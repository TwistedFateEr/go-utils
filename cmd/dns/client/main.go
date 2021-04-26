package main

import (
	"github.com/TwistedFateEr/go-utils/pkg/dnsService"
	"log"
)

func main() {
	r, err := dnsService.QueryIpv4(
		"www.baidu.com",
		"127.0.0.1:9999",
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)
}
