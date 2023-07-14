package main

import (
	"fmt"
	"log"

	client "github.com/kehiy/vv-pactus/client"
	"github.com/kehiy/vv-pactus/utils"
)

var status = make(map[int]string)

type Result struct {
	Address string `json:"adress"`
	Discord string `json:"discord"`
	Status  string `json:"status"`
}

func main() {
	status[1] = "valid"
	status[2] = "offline"
	status[3] = "noSynced"
	status[4] = "invalid"
	result := []Result{}

	data, err := utils.ReadExcel("data.xlsx", "Form Responses 1")
	if err != nil {
		log.Fatalf("error reading data: %v", err)
	}

	fmt.Println(data)
	fmt.Println(status)

	c, err := client.NewClient("172.104.46.145:9090")
	if err != nil {
		log.Fatalf("err making client: %v", err)
	}

	info, err := c.GetNetworkInfo()
	if err != nil {
		log.Fatalf("err read network info: %v", err)
	}
	fmt.Print(info)

	res, err := utils.AddressFromPublicKey("tpublic1p3x4q7a9hu64r2vg27m9a586cvk8nzwxw9raf2ghdfmqqrdsdxktzsgu23kl5p8fsferyje2u8tww6petkydrj3vc8wuhrz8dvldq5z5w4dypxcufvz63rym7d4wjv4n5jvutw8439snva2m89zeqh3kc6qs3f2u9")
	if err != nil {
		log.Fatalf("error drive address from pubkey: %v", err)
	}
	fmt.Println("\n",res)
	fmt.Println(result)
}
