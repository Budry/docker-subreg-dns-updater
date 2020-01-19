package subreg

import (
	"encoding/xml"
	"errors"
	"os"

	"github.com/Budry/subreg-dns-updater-cli/utils"
)

type DNSManager struct {
	Ip string
	Client *SubregCz
}

func (manager *DNSManager) Update(hostname string) error {

	uri, err := utils.NewUri(hostname)
	if err != nil {
		return err
	}

	loginResponse, err := manager.Client.Login(&Login{
		XMLName:  xml.Name{},
		Login:    os.Getenv("SUBREG_USER"),
		Password: os.Getenv("SUBREG_PASSWORD"),
	})

	if err != nil {
		return err
	}

	if loginResponse.Response.Status == "error" {
		return errors.New("Invalid credentials")
	}

	zoneResponse, err := manager.Client.Get_DNS_Zone(&Get_DNS_Zone{
		XMLName: xml.Name{},
		Ssid:    loginResponse.Response.Data.Ssid,
		Domain:  uri.Domain,
	})

	if err != nil {
		return err
	}

	var found *Get_DNS_Zone_Record = nil
	for _, record := range zoneResponse.Response.Data.Records {
		if record.Name == uri.SubDomain {
			found = record
		}
	}

	if found != nil {
		manager.Client.Modify_DNS_Record(&Modify_DNS_Record{
			XMLName: xml.Name{},
			Ssid:    loginResponse.Response.Data.Ssid,
			Domain:  uri.Domain,
			Record: &Modify_DNS_Record_Record{
				Id:      found.Id,
				Type_:   found.Type_,
				Content: manager.Ip,
				Prio:    found.Prio,
				Ttl:     found.Ttl,
			},
		})
	} else {
		manager.Client.Add_DNS_Record(&Add_DNS_Record{
			XMLName: xml.Name{},
			Ssid:    loginResponse.Response.Data.Ssid,
			Domain:  uri.Domain,
			Record: &Add_DNS_Record_Record{
				Name:    uri.SubDomain,
				Type_:   "A",
				Content: manager.Ip,
				Prio:    0,
				Ttl:     0,
			},
		})
	}

	return nil
}