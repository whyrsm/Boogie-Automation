// Package main ...
package main

import (
	"boogie/migrasi"
)

const (
	baseURL            = "http://boogie-nocodb:8080"
	authToken          = "JMAtlF9zwjuKQngMmurdRfC6cHyF4cv1HuSMWZ73"
	customerTableID    = "mc8kwvf7rs295o4"
	sphCustomerTableID = "mmwr1l39roxupyg"
	poCustomerTableID  = "mijyiny797b8vjm"
	articleTableID     = "m58z630zp18936d"
)

func main() {
	migrasi.CustomerAndSPH(baseURL, authToken, customerTableID, sphCustomerTableID)
	// migrasi.CustomerAndPO(baseURL, authToken, customerTableID, poCustomerTableID)
	// migrasi.ArticleAndSPH(baseURL, authToken, articleTableID, sphCustomerTableID)
}
