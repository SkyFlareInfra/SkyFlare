package client

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/miekg/dns"
)

type DNSLookUpInterface interface {
	IsRestrictedDomain(domain string) (bool, error)
	CheckDNSSECStatus(domain string) (bool, error)
	GetNameservers(domain string) ([]string, error)
}

type DNSLookUp struct {
	config pkg.DatabaseConfig
	log    pkg.LogService
}

func NewDNSLookUp(
	config pkg.DatabaseConfig,
	log pkg.LogService,
) DNSLookUpInterface {
	return DNSLookUp{
		config: config,
		log:    log,
	}
}

func (dl DNSLookUp) IsRestrictedDomain(domain string) (bool, error) {
	restrictedDomains := map[string]struct{}{
		"google.com": {}, "facebook.com": {}, "twitter.com": {},
		"amazon.com": {}, "microsoft.com": {}, "cloudflare.com": {},
		"alibaba.com": {}, "ibm.com": {}, "digitalocean.com": {},
		"oracle.com": {}, "linode.com": {}, "vultr.com": {},
		"rackspace.com": {}, "heroku.com": {}, "salesforce.com": {},
		"netflix.com": {}, "uber.com": {}, "airbnb.com": {},
		"stripe.com": {}, "paypal.com": {}, "shopify.com": {},
		"discord.com": {}, "twitch.tv": {}, "github.com": {},
		"reddit.com": {}, "stackoverflow.com": {}, "yelp.com": {},
		"pinterest.com": {}, "snapchat.com": {}, "tiktok.com": {},
		"spotify.com": {}, "etsy.com": {}, "medium.com": {},
		"postman.com": {}, "youtube.com": {}, "instagram.com": {},
		"linkedin.com": {}, "discordapp.com": {}, "slack.com": {},
		"dropbox.com": {}, "soundcloud.com": {}, "flickr.com": {},
		"wordpress.com": {}, "wikipedia.org": {}, "wix.com": {},
		"zillow.com": {}, "zappos.com": {}, "zillowgroup.com": {},
		"zendesk.com": {}, "zendaya.com": {}, "zoho.com": {},
		"zoom.us": {}, "zulily.com": {}, "whatsapp.com": {},
		"weebly.com": {}, "walmart.com": {},
	}
	if _, exists := restrictedDomains[domain]; exists {
		return true, nil
	}

	isPopular := dl.IsPopularDomain(domain)
	if isPopular {
		return true, nil
	}

	// isHosted, err := dl.IsCheckDNS(domain)
	// if err != nil {
	//     dl.logger.Error("Error performing DNS check", err)
	//     return false, err
	// }
	// if isHosted {
	//     return true, nil
	// }

	return false, nil
}

func (dl DNSLookUp) IsPopularDomain(domain string) bool {
	popularSuffixes := []string{
		"google", "facebook", "amazon", "microsoft", "cloudflare",
		"github", "twitter", "instagram", "youtube", "reddit",
		"stackexchange", "stackoverflow", "yelp", "pinterest",
		"discord", "twitch", "etsy", "medium", "spotify",
		"uber", "airbnb", "stripe", "paypal", "shopify",
		"snapchat", "tiktok", "netflix", "apple", "microsoft",
		"amazon", "ebay", "wikipedia", "walmart", "cisco",
	}
	for _, suffix := range popularSuffixes {
		if strings.Contains(domain, suffix) {
			return true
		}
	}
	return false
}

func (dl DNSLookUp) CheckDNSSECStatus(domain string) (bool, error) {
	client := &dns.Client{
		Timeout: 3 * time.Second,
		UDPSize: 4096,
	}

	dnsServers := []string{
		"1.1.1.1:53",
		"8.8.8.8:53",
		"9.9.9.9:53",
		"1.0.0.1:53",
		"8.8.4.4:53",
	}

	checkWithServer := func(server string) (bool, error) {
		msg := new(dns.Msg)
		msg.SetQuestion(dns.Fqdn(domain), dns.TypeDS)
		msg.SetEdns0(4096, true) // Enable DNSSEC OK flag

		resp, _, err := client.Exchange(msg, server)
		if err != nil {
			return false, err
		}

		hasRRSIG := false
		for _, rr := range append(resp.Answer, resp.Ns...) {
			if _, ok := rr.(*dns.RRSIG); ok {
				hasRRSIG = true
				break
			}
		}

		if len(resp.Answer) > 0 || hasRRSIG {
			return true, nil
		}

		msg = new(dns.Msg)
		msg.SetQuestion(dns.Fqdn(domain), dns.TypeDNSKEY)
		resp, _, err = client.Exchange(msg, server)
		if err != nil {
			return false, err
		}

		for _, rr := range resp.Answer {
			if _, ok := rr.(*dns.DNSKEY); ok {
				return true, nil
			}
		}

		return false, nil
	}

	var lastErr error
	for _, server := range dnsServers {
		for i := 0; i < 2; i++ { // 2 tries per server
			enabled, err := checkWithServer(server)
			if err == nil {
				return enabled, nil
			}
			lastErr = err
			time.Sleep(500 * time.Millisecond)
		}
	}

	return false, fmt.Errorf("all DNSSEC checks failed, last error: %w", lastErr)
}

func (dl DNSLookUp) GetNameservers(domain string) ([]string, error) {
	resolvers := []string{
		"8.8.8.8:53",        // Google DNS
		"1.1.1.1:53",        // Cloudflare DNS
		"9.9.9.9:53",        // Quad9 DNS
		"208.67.222.222:53", // OpenDNS
	}

	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	nameservers, err := net.LookupNS(domain)
	if err == nil {
		var nsList []string
		for _, ns := range nameservers {
			nsList = append(nsList, strings.TrimSuffix(ns.Host, "."))
		}
		return nsList, nil
	}

	dl.log.LogError("System DNS lookup failed: %v", err)

	for _, resolver := range resolvers {
		client := &dns.Client{Timeout: 3 * time.Second}
		msg := new(dns.Msg)
		msg.SetQuestion(dns.Fqdn(domain), dns.TypeNS)

		resp, _, err := client.Exchange(msg, resolver)
		if err != nil {
			dl.log.LogErrorFormat("Resolver %s failed: %v", resolver, err)
			continue
		}

		var nsList []string
		for _, rr := range resp.Answer {
			if ns, ok := rr.(*dns.NS); ok {
				nsList = append(nsList, strings.TrimSuffix(ns.Ns, "."))
			}
		}

		if len(nsList) > 0 {
			return nsList, nil
		}
	}

	return nil, fmt.Errorf("all DNS resolvers failed to retrieve nameservers for domain: %s", domain)
}
