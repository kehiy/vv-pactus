package utils

import (
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
