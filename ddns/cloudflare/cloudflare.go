package cloudflare

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"errors"
	"github.com/cloudflare/cloudflare-go"
)

const (
	DEFAULT_ZONE	string	= "meshwalker.com"
	CFEmail		string	= "CF_API_EMAIL"
	CFKey		string	= "CF_API_KEY"
)


var cf *CloudFlare

func New() {
	cf = &CloudFlare{}

	// Prepare CloudFlare credentials
	cfMap := make(map[string]string)
	cfMap[CFEmail]	= os.Getenv(CFEmail)
	cfMap[CFKey]	= os.Getenv(CFKey)

	// Check environment variables
	if err := ValidateConfig(cfMap); err != nil {
		log.Fatal(err)
	}

	// Initialize CloudFlare API
	cf.API = cloudflare.New(cfMap[CFKey], cfMap[CFEmail])

	// Get default zone
	zone_id , err := cf.API.ZoneIDByName(DEFAULT_ZONE)
	if err != nil {
		log.Fatal("Can't find Zone ID for meshwalker.com")
	}
	cf.ZId = zone_id
}



func (cf *CloudFlare) CreateRecord(record NewDNSRecord) error {
	cfRecord := record.Validate()

	err := cf.API.CreateDNSRecord(cf.ZId, *cfRecord)
	if err != nil {
		log.Error("Can't create record", err)
		return err
	}
	return nil
}


func (cf *CloudFlare) UpdateRecord(record NewDNSRecord) error {
	var dnsRecord DNSRecord


	cfRecord, err := cf.API.DNSRecord(cf.ZId, dnsRecord.RecordID)
	if err != nil {
		log.Error("Can't find specified DNS record", err)
		return err
	}

	err = cf.API.UpdateDNSRecord(cf.ZId, dnsRecord.RecordID, cfRecord)
	if err != nil {
		log.Error("Can't update record: ",err)
		return  err
	}

	return nil
}

func (cf *CloudFlare) DeleteRecord(record NewDNSRecord) error {
	return nil
}


func ValidateConfig(m map[string]string) error{
	for key, value := range m {
		if value == "" {
			return errors.New("Environment variable "+key+" is empty or not defined")
		}
	}

	return nil
}