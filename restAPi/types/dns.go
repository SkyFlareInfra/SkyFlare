package types

type DNSRecordInput struct {
	Type         string `json:"type" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Value        string `json:"value" binding:"required"`
	IP           string `json:"IP,omitempty"`
	ProxyEnabled bool   `json:"proxy_enabled,omitempty"`
	ProxyIP      string `json:"proxy_ip,omitempty"`
	OriginIP     string `json:"origin_ip,omitempty"`
	TTL          int    `json:"ttl,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Weight       int    `json:"weight,omitempty"`
	Port         int    `json:"port,omitempty"`
	SerialNumber int    `json:"SerialNumber,omitempty"`
	RefreshTime  int    `json:"RefreshTime,omitempty"`
	RetryTime    int    `json:"RetryTime,omitempty"`
	ExpireTime   int    `json:"ExpireTime,omitempty"`
	MinimumTTL   int    `json:"MinimumTTL,omitempty"`
	Flag         int16  `json:"Flag,omitempty"`
	Flags        string `json:"Flags,omitempty"`
	Tag          string `json:"Tag,omitempty"`
	Usage        int16  `json:"Usage,omitempty"`
	Selector     int16  `json:"Selector,omitempty"`
	MatchingType int16  `json:"MatchingType,omitempty"`
	Certificate  string `json:"Certificate,omitempty"`
	Order        int16  `json:"Order,omitempty"`
	Preference   int16  `json:"Preference,omitempty"`
	Service      string `json:"Service,omitempty"`
	Regexp       string `json:"Regexp,omitempty"`
	Replacement  string `json:"Replacement,omitempty"`
	KeyTag       int16  `json:"KeyTag,omitempty"`
	Algorithm    int16  `json:"Algorithm,omitempty"`
	DigestType   int16  `json:"DigestType,omitempty"`
	Digest       string `json:"Digest,omitempty"`
}

type DNSBulkRecordInput struct {
	Records []DNSRecordInput `json:"records" binding:"required"`
}
