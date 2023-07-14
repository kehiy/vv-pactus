package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	client "github.com/kehiy/vv-pactus/client"
	"github.com/kehiy/vv-pactus/utils"
)

var (
	status             = make(map[int]string)
	validHeight uint32 = 557000
)

type Result struct {
	Address string `json:"adress"`
	Discord string `json:"discord"`
	Status  string `json:"status"`
}

func main() {
	s := time.Now()

	status[1] = "valid"
	status[2] = "invalid"
	status[3] = "notSynced"

	result := []Result{}

	data, err := utils.ReadExcel("data.xlsx", "Form Responses 1")
	if err != nil {
		log.Fatalf("error reading data: %v", err)
	}

	c, err := client.NewClient("94.101.184.118:9090")
	if err != nil {
		log.Fatalf("err making client: %v", err)
	}

	info, err := c.GetNetworkInfo()
	if err != nil {
		log.Fatalf("err read network info: %v", err)
	}

	dup := []string{}

	// check status
	for _, d := range data {
		r := Result{Address: d[1], Discord: d[0]}
		var pub string
	mainl:
		for _, inf := range info.GetPeers() {
			for _, k := range inf.ConsensusKeys {
				pub = utils.AddressFromPublicKey(k)
				if d[1] == pub {
					if inf.Height < validHeight {
						r.Status = status[3]
						result = append(result, r)
						break mainl
					}
					for _, p := range dup {
						if d[1] == p {
							r.Status = status[2]
							result = append(result, r)
							break mainl
						}
					}

					r.Status = status[1]
					result = append(result, r)
					break mainl
				}
				dup = append(dup, pub)
			}
		}
	}

	fin, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("err marshal result: %v", err)
	}

	outputFile, err := os.Create("output.json")
	if err != nil {
		log.Fatalf("err create file: %v", err)
	}
	defer outputFile.Close()

	outputFile.Write([]byte(fin))
	fmt.Print(time.Since(s))
}
