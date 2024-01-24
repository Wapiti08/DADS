package main

import (
	"encoding/xml"
	"log"
	"os"
)

func main() {
	type CrewMember struct {
		ID                int    `xml:"id,omitempty"`
		Name              string `xml:"name"`
		SecurityClearance int    `xml:"clearacelevel"`
		// define the type of sub element
		AccessCodes 	  []string `xml:"accesscodes>code"`
	}

	type ShipInfo struct {
		XMLName   xml.Name `xml:"ship"`
		ShipID    int		`xml:"ShipInfo>ShipID"`
		ShipClass string	`xml:"ShipInfo>ShipClass"`
		Captain   CrewMember
	}

	file, err := os.Create("xmlfile.xml")
	if err != nil {
		log.Fatal("error to create a xml file")
	}
	defer file.Close()

	cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}
	si := ShipInfo{ShipID: 1, ShipClass: "Fighter", Captain: cm}

	// b, err := xml.MarshalIndent(si, " ", "    ")
	// if err != nil {
	// 	fmt.Println("Err:", err)
	// 	return
	// }

	enc := xml.NewEncoder(file)
	enc.Indent(" ","    ")
	enc.Encode(si)

	if err != nil {
		log.Fatal("error on encoding file")
	}
	// fmt.Println(string(b))
}