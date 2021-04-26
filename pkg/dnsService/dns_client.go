package dnsService

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
)

var (
	_dnsClient = new(dnsClient)
)

type dnsClient struct{}

func (this dnsClient) QueryIpFromDns(domain, dnsService string, dnsType uint16) (dst []net.IP, err error) {
	c := dns.Client{}
	m := dns.Msg{}
	m.SetEdns0(4096, true)
	m.SetQuestion(domain+".", dnsType)

	r, rtt, err := c.Exchange(&m, dnsService)
	if err != nil {
		return nil, err
	}

	fmt.Println(r)
	fmt.Println(rtt)

	dst = make([]net.IP, 0, len(r.Answer))
	for _, rr := range r.Answer {
		if i := this.getIpWithDnsType(rr, dnsType); i != nil {
			dst = append(dst, i)
		}
	}
	return
}
func (dnsClient) getIpWithDnsType(rr interface{}, dnstype uint16) net.IP {
	switch dnstype {
	case dns.TypeA:
		v, ok := rr.(*dns.A)
		if ok {
			return v.A
		}
	case dns.TypeAAAA:
		v, ok := rr.(*dns.AAAA)
		if ok {
			return v.AAAA
		}
	}
	return nil
}

func QueryIpv4(domain, dnsService string) ([]net.IP, error) {
	return _dnsClient.QueryIpFromDns(domain, dnsService, dns.TypeA)
}

func QueryIpv6(domain, dnsService string) ([]net.IP, error) {
	return _dnsClient.QueryIpFromDns(domain, dnsService, dns.TypeAAAA)
}
