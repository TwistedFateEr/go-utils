package main

import (
	"fmt"
	"github.com/TwistedFateEr/go-utils/pkg/dnsService"
	"github.com/miekg/dns"
	"log"
)

func main() {
	s := &dnsService.Server{
		Ip:      "127.0.0.1",
		Port:    "9999",
		NetWork: "udp",
	}

	err := s.DnsServer(func(mux *dns.ServeMux) {
		mux.HandleFunc(".", handFunc)
	})
	if err != nil {
		panic(err)
	}
}

func handFunc(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetEdns0(4096, true)
	m.SetReply(r)
	m.Compress = false
	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}
	w.WriteMsg(m)
}

var (
	resip = "2.2.2.2"
	//ipv6  = "2001:4860:4860::6464"
)

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			log.Printf("Query for %s for TypeA\n", q.Name)

			rr, err := dns.NewRR(fmt.Sprintf("%s %d A %s", q.Name, 3600, resip))
			if err == nil {
				log.Printf("%+v\n", rr)
				m.Answer = append(m.Answer, rr)
			}
		}
	}
}
