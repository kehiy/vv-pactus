package main

import (
	"fmt"
	"log"

	"github.com/kehiy/vv-pactus/client"
	"github.com/xuri/excelize/v2"
)

var status = make(map[int]string)

type Result struct {
	Address string
	Discord string
	Status  string
}

func main() {
	status[1] = "valid"
	status[2] = "offline"
	status[3] = "noSynced"
	status[4] = "invalid"

	data, err := readExcel("data.xlsx", "Form Responses 1")
	if err != nil {
		log.Fatalf("error reading data: %v", err)
	}

	fmt.Println(data)
	fmt.Println(status)

	c, err := client.NewClient("172.104.46.145:8080")
	if err != nil {
		log.Fatalf("err making client: %v", err)
	}

	info, err := c.GetNetworkInfo()
	if err != nil {
		log.Fatalf("err read network info: %v", err)
	}
	fmt.Print(info)
}

func readExcel(file string, sheet string) ([][]string, error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
