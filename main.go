package main

import (
	"fmt"
	"log"

	client "github.com/kehiy/vv-pactus/client"
	"github.com/kehiy/vv-pactus/utils"
)

var (
	status = make(map[int]string)
	validHeight uint32 = 557000
) 

type Result struct {
	Address string `json:"adress"`
	Discord string `json:"discord"`
	Status  string `json:"status"`
}

func main() {
	status[1] = "valid"
	status[2] = "invalid"
	status[3] = "notSynced"

	result := []Result{}

	data, err := utils.ReadExcel("data.xlsx", "Form Responses 1")
	if err != nil {
		log.Fatalf("error reading data: %v", err)
	}

	c, err := client.NewClient("172.104.46.145:9090")
	if err != nil {
		log.Fatalf("err making client: %v", err)
	}

	info, err := c.GetNetworkInfo()
	if err != nil {
		log.Fatalf("err read network info: %v", err)
	}

	// check notSynced status
	for _, d := range data {
		r := Result{Address: d[1], Discord: d[0]}
		for _, inf := range info.GetPeers() {
			for _, k := range inf.ConsensusKeys {
				if d[1] == utils.AddressFromPublicKey(k) {
					if inf.Height < validHeight {
						r.Status = status[3]
						result = append(result, r)
					}
				}
			}
		}
	}

	fmt.Println(result)
}
