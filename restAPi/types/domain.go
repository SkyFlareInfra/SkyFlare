package types

type DomainInput struct {
	Domain         string           `json:"domain" binding:"required"`
	DNSRecords     []DNSRecordInput `json:"dns_records,omitempty"`
	ZoneFile       string           `json:"zone_file,omitempty"`
}
