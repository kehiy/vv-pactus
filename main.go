package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

var status = make(map[int]string)

func main() {
	status[1] = "valid"
	status[2] = "offline"
	status[3] = "noSynced"
	status[4] = "invalid"

	data, err := readExcel("data.xlsx", "Form Responses 1")
	if err != nil {
		fmt.Print("error")
	}
	fmt.Println(data)
	fmt.Println(status)
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
