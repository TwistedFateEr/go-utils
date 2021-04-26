package netlinkService

import "github.com/vishvananda/netlink"

func AddAddr(linkname, addrinfo string) error {
	lo, err := netlink.LinkByName(linkname)
	if err != nil {
		return err
	}

	addr, err := netlink.ParseAddr(addrinfo)
	if err != nil {
		return err
	}
	return netlink.AddrReplace(lo, addr)
}

func DelAddr(linkname, addrinfo string) error {
	lo, err := netlink.LinkByName(linkname)
	if err != nil {
		return err
	}

	addr, err := netlink.ParseAddr(addrinfo)
	if err != nil {
		return err
	}
	return netlink.AddrDel(lo, addr)
}


func FindAllAddr(linkname string) ([]netlink.Addr, error) {
	ln, err := netlink.LinkByName(linkname)
	if err != nil {
		return nil, err
	}
	return netlink.AddrList(ln, netlink.NDA_LLADDR)
}


