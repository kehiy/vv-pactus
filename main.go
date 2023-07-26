package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	client "github.com/kehiy/vv-pactus/client"
	"github.com/kehiy/vv-pactus/utils"
	"github.com/libp2p/go-libp2p/core/peer"
)

var (
	validHeight uint32 = 600000
)

type Result struct {
	ID            int    `json:"id"`
	Address       string `json:"address"`
	Discord       string `json:"discord"`
	DiscordHide   string `json:"discordhide"`
	Instagram     string `json:"instagram"`
	InstagramHide string `json:"instagramhide"`
	Twitter       string `json:"twitter"`
	TwitterHide   string `json:"twitterhide"`
	Status        string `json:"status"`
	PeerId        string `json:"peerid"`
	ValNum        int32  `json:"validaornumber"`
	ValSeq        int32  `json:"validaorseq"`
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
		if len(d) == 1 {
			continue
		}
		r := Result{Address: d[3], Discord: d[0], Twitter: d[1], Instagram: d[2]}
		var addr string
		for _, inf := range info.GetPeers() {
			for _, k := range inf.ConsensusKeys {
				addr = utils.AddressFromPublicKey(k)
				if d[1] == addr {
					status := "valid"
					if inf.Height < validHeight {
						status = "notSynced"
					}
					index, ok := dup[string(inf.GetPeerId())]
					if ok {
						status = "duplicate"
						result[index].Status = "duplicate"
					} else {
						dup[string(inf.GetPeerId())] = len(result)
					}

					r.Status = status
					pid, _ := peer.IDFromBytes(inf.GetPeerId())
					r.PeerId = pid.String()
					r.ID = len(result) + 1
					r.DiscordHide = utils.HideId(r.Discord)
					r.InstagramHide = utils.HideId(r.Instagram)
					r.TwitterHide = utils.HideId(r.Twitter)
					validatorInfo, err := c.GetValidatorInfo(r.Address)
					if err == nil {
						r.ValNum = validatorInfo.Validator.Number
						r.ValSeq = validatorInfo.Validator.Sequence
					}
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
