package dnsService

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
)

type Server struct {
	Ip      string
	Port    string
	NetWork string

	ExitCTX context.Context
}

func (s Server) DnsServer(apis ...func(mux *dns.ServeMux)) error {
	sux := dns.NewServeMux()

	for _, api := range apis {
		api(sux)
	}

	server := &dns.Server{
		Addr:    fmt.Sprintf("%v:%v", s.Ip, s.Port),
		Net:     s.NetWork,
		Handler: sux,
	}
	defer server.Shutdown()

	return server.ListenAndServe()
}
