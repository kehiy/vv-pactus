package utils

import (
	"strings"

	"github.com/pactus-project/pactus/crypto"
	pactus "github.com/pactus-project/pactus/crypto/bls"
	"github.com/xuri/excelize/v2"
)

func ReadExcel(file string, sheet string) ([][]string, error) {
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

func AddressFromPublicKey(s string) string {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	pub, _ := pactus.PublicKeyFromString(s)
	addr := pub.Address().String()
	return addr
}

func Search(slice []string, target string) string {
	for _, value := range slice {
		if value == target {
			return value
		}
	}
	return ""
}

func IsIn(s string, slice []string) bool {
	for _, d := range slice {
		if s == d {
			return true
		}
	}
	return false
}

func HideId(input string) string {
	if len(input) < 2 {
		return input
	}

	first := string(input[0])
	last := string(input[len(input)-1])
	stars := strings.Repeat("*", 10)

	return first + stars + last
}
