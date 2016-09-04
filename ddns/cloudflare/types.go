package cloudflare

import (
	"time"
	"github.com/cloudflare/cloudflare-go"
)

type DNSProvider interface {
	Create(NewDNSRecord) error
	Update(NewDNSRecord) error
	Delete(NewDNSRecord) error
}

type CloudFlare struct {
	API         	*cloudflare.API
	ZId        	string
}

type DNSRecord struct {
	RecordID	string	`json:"record_id,omitempty"`
	Name		string	`json:"name,omitempty"`
	Content		string	`json:"content,omitempty"`
	Type		string	`json:"type,omitempty"`
	TTL		int	`json:"ttl,omitempty"`
	Domain		string	`json:"domain,omitempty"`
	Priority	int	`json:"priority"`
}

type NewDNSRecord struct {
	Name		string	`json:"name,omitempty"`
	Domain		string	`json:"domain,omitempty"`
	Content		string	`json:"content,omitempty"`
	Type		string	`json:"type,omitempty"`
	TTL		int	`json:"ttl,omitempty"`
}

type Zone struct {
	ID		int		`json:"id,omitempty"`
	ZoneId		string		`json:"cf_zone_id"`
	Name		string		`json:"name,omitempty"`
	CreatedAt	time.Time	`json:"created_on"`
	ModifiedAt	time.Time	`json:"modified_on"`
}

type Cluster struct {
	ID	int	`json:"id,omitempty"`
	Token	string	`json:"token"`
}

type ErrorMsg struct {
	ErrorCode	int	`json:"errorcode,omitempty"`
	Message		string	`json:"message,omitempty"`
}

