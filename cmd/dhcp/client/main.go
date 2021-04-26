package main

import (
	"flag"
	"fmt"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/client4"
	"github.com/insomniacslk/dhcp/dhcpv6"
	"github.com/insomniacslk/dhcp/dhcpv6/client6"
	"github.com/insomniacslk/dhcp/netboot"
	"log"
	"net"
)

var (
	ifname  = flag.String("i", "ens33", "Interface name")
)

func dhclient6(ifname string, verbose bool) (error) {
	llAddr, err := dhcpv6.GetLinkLocalAddr(ifname)
	if err != nil {
		return err
	}
	laddr := net.UDPAddr{
		IP:   llAddr,
		Port: dhcpv6.DefaultClientPort,
		Zone: ifname,
	}
	raddr := net.UDPAddr{
		IP:   dhcpv6.AllDHCPRelayAgentsAndServers,
		Port: dhcpv6.DefaultServerPort,
		Zone: ifname,
	}
	c := client6.NewClient()
	c.LocalAddr = &laddr
	c.RemoteAddr = &raddr
	var conv []dhcpv6.DHCPv6

	i := 0
	for{
		conv, err = c.Exchange(ifname, dhcpv6.WithNetboot)
		if err == nil{
			break
		}
		if i >=5 {
			return fmt.Errorf("i :%v",i)
		}

		i++
		log.Printf("Exchange: %v", err)
	}
	if verbose {
		for _, m := range conv {
			log.Printf("%+v\n", m.Summary())
		}
	}

	// extract the network configuration
	netconf, err := netboot.ConversationToNetconf(conv)
	if err != nil {
		return err
	}

	return netboot.ConfigureInterface(ifname, &netconf.NetConf)
}

func dhclient4(ifname string, verbose bool) (error) {
	client := client4.NewClient()
	var (
		conv []*dhcpv4.DHCPv4
		err  error
	)

	for{
		conv, err = client.Exchange(ifname)
		if err == nil {
			break
		}
		log.Printf("Exchange  err:%v",err)
	}


	if verbose {
		for _, m := range conv {
			log.Printf("%+v\n", m)
		}
	}

	// extract the network configuration
	netconf, err := netboot.ConversationToNetconfv4(conv)
	if err != nil {
		return err
	}

	return netboot.ConfigureInterface(ifname, &netconf.NetConf)
}

func main() {
	flag.Parse()

	var  err error

	err = dhclient4(*ifname,true)
	if err != nil {
		log.Fatal(err)
	}

	err= dhclient6(*ifname,true)
	if err != nil {
		log.Fatal(err)
	}
}
