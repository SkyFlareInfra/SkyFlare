package utility

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/SkyFlareInfra/SkyFlare/models"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/types"
	"github.com/jellydator/ttlcache/v3"
	"github.com/miekg/dns"
)

type DNSUtilityInterface interface {
	ParseZoneFile(zoneFile string) ([]models.DNSRecord, error)
	ScanDNSRecords(domain string) ([]models.DNSRecord, error)
	CreateDNSRecordsFromInput(inputs []types.DNSRecordInput) []models.DNSRecord
	GetZoneFile(domain string) (string, error)
	GetDNSResponse(domain string, recordType uint16) ([]models.DNSRecord, error)
}

type DNSUtility struct {
	log     pkg.LogService
	cache   *ttlcache.Cache[string, []models.DNSRecord]
	errorMu sync.Mutex
	config  pkg.DatabaseConfig
}

func NewDNSLookUp(
	log pkg.LogService,
	config pkg.DatabaseConfig,

) DNSUtilityInterface {
	cache := ttlcache.New(
		ttlcache.WithTTL[string, []models.DNSRecord](5 * time.Minute),
	)

	return &DNSUtility{
		log:     log,
		cache:   cache,
		config:  config,
		errorMu: sync.Mutex{},
	}
}

// ParseZoneFile parses a DNS zone file and returns DNS records
func (du *DNSUtility) ParseZoneFile(zoneFile string) ([]models.DNSRecord, error) {
	var records []models.DNSRecord

	scanner := bufio.NewScanner(strings.NewReader(zoneFile))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}

		zoneRecord, err := dns.NewRR(line)
		if err != nil {
			du.log.LogErrorFormat("failed to parse zone record: %v", err)
			return nil, fmt.Errorf("failed to parse zone record: %w", err)
		}

		recordType, ok := dns.TypeToString[zoneRecord.Header().Rrtype]
		if !ok {
			du.log.LogErrorFormat("Unknown DNS record type: %d. Using default value 'Unknown'.", zoneRecord.Header().Rrtype)

			recordType = "Unknown"
		}

		dnsRecord := models.DNSRecord{
			Type:  recordType, // Record type (A, CNAME, etc.)
			Name:  zoneRecord.Header().Name,
			Value: zoneRecord.String(),
			TTL:   int(zoneRecord.Header().Ttl),
		}
		records = append(records, dnsRecord)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

// ScanDNSRecords queries DNS records for a domain
func (du *DNSUtility) ScanDNSRecords(domain string) ([]models.DNSRecord, error) {
	// Check cache first
	if item := du.cache.Get(domain); item != nil {
		return item.Value(), nil
	}

	var records []models.DNSRecord
	var mu sync.Mutex

	resolvers := []string{
		"1.1.1.1:53",        // Cloudflare
		"8.8.8.8:53",        // Google
		"9.9.9.9:53",        // Quad9
		"208.67.222.222:53", // OpenDNS
	}

	recordTypes := []uint16{
		dns.TypeA,
		dns.TypeAAAA,
		dns.TypeMX,
		dns.TypeNS,
		dns.TypeTXT,
		dns.TypeCNAME,
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 8) // Limit concurrent queries

	for _, resolver := range resolvers {
		for _, recordType := range recordTypes {
			wg.Add(1)
			semaphore <- struct{}{}

			go func(resolver string, recordType uint16) {
				defer wg.Done()
				defer func() { <-semaphore }()

				client := &dns.Client{Timeout: 2 * time.Second}
				m := new(dns.Msg)
				m.SetQuestion(dns.Fqdn(domain), recordType)
				m.RecursionDesired = true

				resp, _, err := client.Exchange(m, resolver)
				if err != nil {
					du.log.LogErrorFormat("DNS query failed for %s type %d via %s: %v",
						domain, recordType, resolver, err)
					return
				}

				if resp == nil || resp.Rcode != dns.RcodeSuccess {
					return
				}

				newRecords := parseDNSAnswers(domain, resp.Answer)
				if len(newRecords) > 0 {
					mu.Lock()
					records = append(records, newRecords...)
					mu.Unlock()
				}
			}(resolver, recordType)
		}
	}

	wg.Wait()

	// Try system resolvers if no records found
	if len(records) == 0 {
		conf, err := dns.ClientConfigFromFile("/etc/resolv.conf")
		if err == nil && len(conf.Servers) > 0 {
			for _, recordType := range recordTypes {
				client := &dns.Client{Timeout: 2 * time.Second}
				m := new(dns.Msg)
				m.SetQuestion(dns.Fqdn(domain), recordType)
				m.RecursionDesired = true

				for _, resolver := range conf.Servers {
					resp, _, err := client.Exchange(m, resolver+":"+conf.Port)
					if err != nil {
						continue
					}

					if resp != nil && resp.Rcode == dns.RcodeSuccess {
						newRecords := parseDNSAnswers(domain, resp.Answer)
						if len(newRecords) > 0 {
							records = append(records, newRecords...)
						}
					}
				}
			}
		}
	}

	deduped := deduplicateRecords(records)

	if len(deduped) > 0 {
		du.cache.Set(domain, deduped, ttlcache.DefaultTTL)
	}

	// Return empty slice instead of error if no records found
	return deduped, nil
}

func parseDNSAnswers(domain string, answers []dns.RR) []models.DNSRecord {
	var records []models.DNSRecord

	for _, ans := range answers {
		switch record := ans.(type) {
		// Essential Records
		case *dns.A:
			records = append(records, models.DNSRecord{
				Type:  "A",
				Name:  domain,
				Value: record.A.String(),
				IP:    record.A.String(),
				TTL:   int(record.Hdr.Ttl),
			})

		case *dns.AAAA:
			records = append(records, models.DNSRecord{
				Type:  "AAAA",
				Name:  domain,
				Value: record.AAAA.String(),
				IP:    record.AAAA.String(),
				TTL:   int(record.Hdr.Ttl),
			})

		case *dns.MX:
			records = append(records, models.DNSRecord{
				Type:     "MX",
				Name:     domain,
				Value:    record.Mx,
				TTL:      int(record.Hdr.Ttl),
				Priority: int(record.Preference),
			})

		case *dns.NS:
			records = append(records, models.DNSRecord{
				Type:  "NS",
				Name:  domain,
				Value: record.Ns,
				TTL:   int(record.Hdr.Ttl),
			})

		case *dns.CNAME:
			records = append(records, models.DNSRecord{
				Type:  "CNAME",
				Name:  domain,
				Value: record.Target,
				TTL:   int(record.Hdr.Ttl),
			})

		case *dns.TXT:
			records = append(records, models.DNSRecord{
				Type:  "TXT",
				Name:  domain,
				Value: strings.Join(record.Txt, " "),
				TTL:   int(record.Hdr.Ttl),
			})

		case *dns.SOA:
			records = append(records, models.DNSRecord{
				Type:         "SOA",
				Name:         domain,
				Value:        fmt.Sprintf("Primary NS: %s, Admin Email: %s", record.Ns, record.Mbox),
				TTL:          int(record.Hdr.Ttl),
				SerialNumber: int(record.Serial),
				RefreshTime:  int(record.Refresh),
				RetryTime:    int(record.Retry),
				ExpireTime:   int(record.Expire),
				MinimumTTL:   int(record.Minttl),
			})

		// Security Records
		case *dns.CAA:
			records = append(records, models.DNSRecord{
				Type:  "CAA",
				Name:  domain,
				Tag:   record.Tag,
				Value: record.Value,
				Flag:  int16(record.Flag),
				TTL:   int(record.Hdr.Ttl),
			})

		case *dns.TLSA:
			records = append(records, models.DNSRecord{
				Type:         "TLSA",
				Name:         domain,
				Usage:        int16(record.Usage),
				Selector:     int16(record.Selector),
				MatchingType: int16(record.MatchingType),
				Certificate:  record.Certificate,
				TTL:          int(record.Hdr.Ttl),
			})

		// Service Records
		case *dns.SRV:
			records = append(records, models.DNSRecord{
				Type:     "SRV",
				Name:     domain,
				Value:    record.Target,
				TTL:      int(record.Hdr.Ttl),
				Priority: int(record.Priority),
				Weight:   int(record.Weight),
				Port:     int(record.Port),
			})

		case *dns.NAPTR:
			records = append(records, models.DNSRecord{
				Type:        "NAPTR",
				Name:        domain,
				Order:       int16(record.Order),
				Preference:  int16(record.Preference),
				Flags:       record.Flags,
				Service:     record.Service,
				Regexp:      record.Regexp,
				Replacement: record.Replacement,
				TTL:         int(record.Hdr.Ttl),
			})

		// DNSSEC Records
		case *dns.DS:
			records = append(records, models.DNSRecord{
				Type:       "DS",
				Name:       domain,
				KeyTag:     int16(record.KeyTag),
				Algorithm:  int16(record.Algorithm),
				DigestType: int16(record.DigestType),
				Digest:     record.Digest,
				TTL:        int(record.Hdr.Ttl),
			})

		// Special Cases
		case *dns.PTR:
			records = append(records, models.DNSRecord{
				Type:  "PTR",
				Name:  domain,
				Value: record.Ptr,
				TTL:   int(record.Hdr.Ttl),
			})

		case *dns.HINFO:
			records = append(records, models.DNSRecord{
				Type:  "HINFO",
				Name:  domain,
				Value: fmt.Sprintf("%s %s", record.Cpu, record.Os),
				TTL:   int(record.Hdr.Ttl),
			})

		default:
			// Skip unsupported record types
			continue
		}
	}

	return records
}

// deduplicateRecords removes duplicate DNS records
func deduplicateRecords(records []models.DNSRecord) []models.DNSRecord {
	seen := make(map[string]struct{})
	var unique []models.DNSRecord

	for _, r := range records {
		key := fmt.Sprintf("%s|%s|%s|%d", r.Type, r.Name, r.Value, r.TTL)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			unique = append(unique, r)
		}
	}

	return unique
}

// Utility to create DNS records from input
func (du *DNSUtility) CreateDNSRecordsFromInput(inputs []types.DNSRecordInput) []models.DNSRecord {
	records := []models.DNSRecord{}
	for _, input := range inputs {
		record := models.DNSRecord{
			Type:         input.Type,
			Name:         input.Name,
			Value:        input.Value,
			TTL:          input.TTL,
			Priority:     input.Priority,
			Weight:       input.Weight,
			SerialNumber: input.SerialNumber,
			RefreshTime:  input.RefreshTime,
			RetryTime:    input.RetryTime,
			ExpireTime:   input.ExpireTime,
			MinimumTTL:   input.MinimumTTL,
			Flag:         input.Flag,
			Flags:        input.Flags,
			Tag:          input.Tag,
			Usage:        input.Usage,
			Selector:     input.Selector,
			MatchingType: input.MatchingType,
			Certificate:  input.Certificate,
			Order:        input.Order,
			Preference:   input.Preference,
			Service:      input.Service,
			Regexp:       input.Regexp,
			Replacement:  input.Replacement,
			KeyTag:       input.KeyTag,
			Algorithm:    input.Algorithm,
			DigestType:   input.DigestType,
			Digest:       input.Digest,
		}

		if input.Type == "SRV" {
			record.Weight = input.Weight
			record.Port = input.Port
		}

		records = append(records, record)
	}
	return records
}

func (du *DNSUtility) GetZoneFile(domain string) (string, error) {
	cmd := fmt.Sprintf("dig +short %s AXFR", domain)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil || strings.TrimSpace(string(output)) == "" {
		return "", nil
	}

	return string(output), nil
}

func (du *DNSUtility) GetDNSResponse(domain string, recordType uint16) ([]models.DNSRecord, error) {
	records, err := du.ScanDNSRecords(domain)
	if err != nil {
		return nil, err
	}

	for i := range records {
		if records[i].ProxyEnabled {
			switch records[i].Type {
			case "A":
				if records[i].ProxyIP != "" {
					records[i].Value = records[i].ProxyIP
					records[i].IP = records[i].ProxyIP
				} else if du.config.ProxyIPv4 != "" {
					records[i].Value = du.config.ProxyIPv4
					records[i].IP = du.config.ProxyIPv4
					records[i].ProxyIP = du.config.ProxyIPv4
				}

				if records[i].OriginIP == "" {
					records[i].OriginIP = records[i].Value
				}

			case "AAAA":
				if records[i].ProxyIP != "" {
					records[i].Value = records[i].ProxyIP
					records[i].IP = records[i].ProxyIP
				} else if du.config.ProxyIPv6 != "" {
					records[i].Value = du.config.ProxyIPv6
					records[i].IP = du.config.ProxyIPv6
					records[i].ProxyIP = du.config.ProxyIPv6
				}

				if records[i].OriginIP == "" {
					records[i].OriginIP = records[i].Value
				}
			}
		} else {
			switch records[i].Type {
			case "A", "AAAA":
				records[i].Value = records[i].OriginIP
				records[i].IP = records[i].OriginIP
				records[i].ProxyIP = ""
			}
		}
	}

	return records, nil
}
