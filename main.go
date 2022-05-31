package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fawkesley/pollution-map/datastore"
	"github.com/fawkesley/pollution-printouts/addresspollution"
)

var postcodes = []string{
	"L15 0EB",
	"L15 2HD",
}

func main() {

	email := os.Getenv("USER_AGENT_EMAIL")
	if email == "" {
		log.Panic("please set a contact email in USER_AGENT_EMAIL")
	}

	ap, err := addresspollution.NewClient(
		email,
		"https://github.com/fawkesley/pollution-printouts",
	)
	if err != nil {
		log.Panic(err)
	}

	loadTargetPostcodes()

	var addresses []addresspollution.Address

	for _, pc := range postcodes {
		fmt.Println(pc)
		a, err := ap.Addresses(pc)
		if err != nil {
			log.Panic(err)
		}

		addresses = append(addresses, a...)
	}

	for _, addr := range addresses {
		rating, err := ap.PollutionAtAddress(addr.ID)
		if err != nil {
			log.Panic(err)
		}

		err = datastore.SaveAddress(addr.ID, *rating)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("saved %s\n", addr.FormattedAddress)
		//datastore.SaveAddress(
		//	uuid.Must(uuid.FromString("634f544a-781e-4ad9-b17f-cb72d5b4933c")),
		//	addresspollution.PollutionLevels{
		//		Pm2_5: 10.5,
		//		Pm10:  15.5,
		//		No2:   20.5,
		//	},
		//)

		// fmt.Printf("⚠️ %s AIR POLLUTION ⚠️\n", strings.ToUpper(rating.PollutionDescription))
		// fmt.Printf("%s\n\n", rating.FormattedAddress)
		// fmt.Printf("At your home, %d out of 3 pollutants\n", rating.NumPollutantsExceedingLimits())
		// fmt.Printf("exceed World Health Organisation safe levels.\n\n")

		// fmt.Printf("PM2.5\n")
		// fmt.Printf("%.1f\n", rating.Pm2_5)
		// fmt.Printf("%s safe level\n\n", rating.Pm2_5SafeLevelDescription())

		// fmt.Printf("PM10\n")
		// fmt.Printf("%.1f\n", rating.Pm10)
		// fmt.Printf("%s safe level\n\n", rating.Pm10SafeLevelDescription())

		// fmt.Printf("NO2\n")
		// fmt.Printf("%.1f\n", rating.No2)
		// fmt.Printf("%s safe level\n\n", rating.No2SafeLevelDescription())
	}

}

// loadTargetPostcodes opens target_postcodes.txt and loads each line into `postcodes`
func loadTargetPostcodes() {
}

func slugify(address string) string {
	address = strings.ToLower(address)
	address = strings.Replace(address, " ", "-", -1)
	address = strings.Replace(address, ",", "", -1)
	return address

}
