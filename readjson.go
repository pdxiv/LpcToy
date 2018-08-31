package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Segment struct {
	Energy float32   `json:"Energy"`
	Period float32   `json:"Period"`
	K      []float32 `json:"K"`
}

type PhoneDataFormat map[string][]Segment

func main() {
	phone := loadJSONData("phones.json")
	for phoneSymbol, phoneData := range phone {
		fmt.Println(phoneSymbol)
		fmt.Println("  Period", phoneData[0].Period)
		fmt.Println("  Energy", phoneData[0].Energy)
		for i, k := range phoneData[0].K {
			fmt.Println("  K", i, k)
		}
	}

	if found, _ := findSymbol(phone, "AA"); found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}

// Check if a phone symbol exists in the phone database. if so, return it
func findSymbol(phone PhoneDataFormat, toFind string) (bool, []Segment) {
	if val, ok := phone[toFind]; ok {
		return true, val
	} else {
		return false, nil
	}
}

func loadJSONData(filename string) PhoneDataFormat {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(strings.NewReader(string(content)))
	var phone PhoneDataFormat
	for {
		if err := dec.Decode(&phone); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return phone
}
