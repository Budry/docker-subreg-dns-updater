package cmd

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Budry/subreg-dns-updater-cli/subreg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo [DOMAIN] [IP] [DNS NAME]",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.   
                Complete documentation is available at http://hugo.spf13.com`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		domain := args[0]
		ip := args[1]
		name := args[2]

		service := subreg.NewSubregCz("", false, &subreg.BasicAuth{})

		loginResponse, err := service.Login(&subreg.Login{
			XMLName:  xml.Name{},
			Login:    os.Getenv("SUBREG_USER"),
			Password: os.Getenv("SUBREG_PASSWORD"),
		})

		if err != nil {
			panic(err)
		}

		zoneResponse, err := service.Get_DNS_Zone(&subreg.Get_DNS_Zone{
			XMLName: xml.Name{},
			Ssid:    loginResponse.Response.Data.Ssid,
			Domain:  domain,
		})

		var found *subreg.Get_DNS_Zone_Record = nil
		for _, record := range zoneResponse.Response.Data.Records {
			if record.Name == name {
				found = record
			}
		}
		if found != nil {
			service.Modify_DNS_Record(&subreg.Modify_DNS_Record{
				XMLName: xml.Name{},
				Ssid:    loginResponse.Response.Data.Ssid,
				Domain:  domain,
				Record: &subreg.Modify_DNS_Record_Record{
					Id:      found.Id,
					Type_:   found.Type_,
					Content: ip,
					Prio:    found.Prio,
					Ttl:     found.Ttl,
				},
			})
		} else {
			service.Add_DNS_Record(&subreg.Add_DNS_Record{
				XMLName: xml.Name{},
				Ssid:    loginResponse.Response.Data.Ssid,
				Domain:  domain,
				Record: &subreg.Add_DNS_Record_Record{
					Name:    name,
					Type_:   "A",
					Content: ip,
					Prio:    0,
					Ttl:     0,
				},
			})
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
