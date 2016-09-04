package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
)


func (r *NewDNSRecord) Validate() *cloudflare.DNSRecord{
	switch r.Type {
	case("A"):
		return &cloudflare.DNSRecord{
			Type:		r.Type,
			Name:		r.Name,
			Content:	r.Content,
			TTL:		r.TTL,
			Proxiable:	false,
			Proxied:	false,
		}

	case("AAAA"):
		return &cloudflare.DNSRecord{
			Type:		r.Type,
			Name:		r.Name,
			Content:	r.Content,
			TTL:		r.TTL,
			Proxiable:	false,
			Proxied:	false,
		}

	case("CNAME"):
		return &cloudflare.DNSRecord{
			Type:		r.Type,
			Name:		r.Name,
			Content:	r.Content,
			TTL:		r.TTL,
			Proxiable:	false,
			Proxied:	false,
		}

	case("MX"):
		return &cloudflare.DNSRecord{
			Type:		r.Type,
			Name:		r.Name,
			Content:	r.Content,
			TTL:		r.TTL,
			Proxiable:	false,
			Proxied:	false,
		}
	}

	return nil
}


/*
func (record *NewDNSRecord) ValidateARecord() error{
	ip := net.ParseIP(record.Content)
	if ip == nil {
		return errors.New("No valid ip address found")
	}

	ipv4 := ip.To4()
	if ipv4 == nil {
		return errors.New("No valid ip address found")
	}

	return nil
}


func (record *NewDNSRecord) ValidateAAAARecord() error{
	ip := net.ParseIP(record.Content)
	if ip == nil {
		return errors.New("No valid IPv4 address found")
	}

	ipv6 := ip.To16()
	if ipv6 == nil {
		return errors.New("No valid IPv6 address found")
	}

	return nil
}


func (ndr *NewDNSRecord) ValidateCNAMERecord() error{
	return nil
}


func (ndr *NewDNSRecord) ValidateMXRecord() error{
	return nil
}
*/