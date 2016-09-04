package ddns

import (
	log "github.com/Sirupsen/logrus"
	cf "meshwalker.com/mws/ddns/cloudflare"
	"meshwalker.com/mws/cluster"
	"os"
)

func Init() {
	cf.New()	// Init CloudFlare backend (DNS)
}

const (
	CLUSTER_IP	= os.Getenv("CLUSTER_IP")

)

func SetFailureRecords(cluster *cluster.Cluster) {
	aRecord := &cf.NewDNSRecord{
		Type:		"A",
		Name:		CLUSTER_IP,
		Content:	cluster.FullDomain,
		TTL:		60,
	}

	if err := cf.DNSProvider.Create(aRecord); err != nil {
		log.Error(err)
	}

	// Build failure redirect url
	rUrl := cf.DEFAULT_ZONE+"/info/"+cluster.Name
	/*
	cnameRecord := &cf.NewDNSRecord{
		Type:		"CNAME",
		Name:		cf.DEFAULT_ZONE,
		Content:	rUrl,
		TTL:		60,
	}

	if err := cf.DNSProvider.Create(cnameRecord); err != nil {
		log.Error(err)
	}*/
}

func SetClusterRecords(cluster *cluster.Cluster) {
	aRecord := &cf.NewDNSRecord{
		Type:		"A",
		Name:		cluster.IP.String(),
		Content:	cluster.FullDomain,
		TTL:		60,
	}

	if err := cf.DNSProvider.Create(aRecord); err != nil {
		log.Error(err)

	}

	cnameRecord := &cf.NewDNSRecord{
		Type:		"CNAME",
		Name:		cluster.FullDomain,
		Content:	"www",
		TTL:		60,
	}

	if err := cf.DNSProvider.Create(cnameRecord); err != nil {
		log.Error(err)
	}
}