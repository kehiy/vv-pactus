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
	validHeight uint32 = 557000
)

type Result struct {
	Address string `json:"adress"`
	Discord string `json:"discord"`
	Status  string `json:"status"`
	PeerId  string `json:"peerid"`
}

func main() {
	s := time.Now()

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

	dup := make(map[string]int)

	// check status
	for _, d := range data {
		r := Result{Address: d[1], Discord: d[0]}
		var addr string
	mainl:
		for _, inf := range info.GetPeers() {
			for _, k := range inf.ConsensusKeys {
				addr = utils.AddressFromPublicKey(k)
				if d[1] == addr {
					status := "valid"
					if inf.Height < validHeight {
						status = "notSynced"
						break mainl
					}
					index, ok := dup[string(inf.GetPeerId())]
					if ok {
						status = "duplicate"
						result[index].Status = "duplicate"
						break mainl
					}
					
					dup[string(inf.GetPeerId())] = len(result)
					r.Status = status
					r.PeerId = string(inf.GetPeerId())
					result = append(result, r)
				}
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
