package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type Segment struct {
	Energy float32   `json:"Energy"`
	Period float32   `json:"Period"`
	K      []float32 `json:"K"`
}

type Person struct {
	// Name string `json:"name"`
	// Age  int    `json:"age"`

}

type Info map[string][]Segment

const jsonStream = `
{
	"AA": [{
			"Energy": 81,
			"K": [-24256,
				9216,
				65,
				1,
				34, -20, -41,
				55,
				37, -15
			],
			"Period": 54
		},
		{
			"Energy": 81,
			"K": [-21632,
				6720,
				65,
				57,
				11,
				36, -53,
				55,
				57,
				22
			],
			"Period": 69
		},
		{
			"Energy": 81,
			"K": [-21632,
				1536,
				79,
				57, -24,
				47, -41,
				79,
				57,
				22
			],
			"Period": 71
		},
		{
			"Energy": 81,
			"K": [-21632,
				1536,
				79,
				57, -24,
				47, -41,
				79,
				57,
				22
			],
			"Period": 66
		},
		{
			"Energy": 81,
			"K": [-24256,
				19648,
				25, -26,
				46,
				25,
				7,
				31,
				57, -33
			],
			"Period": 69
		}
	],
	"AE": [{
			"Energy": 57,
			"K": [-10048,
				6720,
				38, -12, -1,
				58,
				19,
				102,
				16,
				4
			],
			"Period": 99
		},
		{
			"Energy": 57,
			"K": [-5184,
				9216,
				38, -12, -12,
				58,
				7,
				102,
				37,
				22
			],
			"Period": 95
		},
		{
			"Energy": 114,
			"K": [-5184,
				13824,
				52, -26, -1,
				47,
				7,
				102,
				37,
				4
			],
			"Period": 92
		},
		{
			"Energy": 81,
			"K": [-10048,
				26176,
				11, -54,
				34, -20,
				19,
				79, -4,
				4
			],
			"Period": 92
		},
		{
			"Energy": 2,
			"K": [
				10048, -3712,
				25, -12, -24, -20, -53,
				55, -4, -51
			],
			"Period": 16
		},
		{
			"Energy": 0,
			"K": [
				10048, -3712,
				25, -12, -24, -20, -53,
				55, -4, -51
			],
			"Period": 16
		},
		{
			"Energy": 0,
			"K": [
				10048, -3712,
				25, -12, -24, -20, -53,
				55, -4, -51
			],
			"Period": 16
		}
	]
}`

func main() {
	fmt.Println("Starting...")
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var info Info
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(info["AA"][0].Energy)
		fmt.Println(info["AE"][0].Energy)

	}
}
