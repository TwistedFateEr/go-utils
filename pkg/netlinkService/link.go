package netlinkService

import "github.com/vishvananda/netlink"

func AddLink(linkname string,fs ...func(attrs netlink.LinkAttrs)) error {
	_, err := netlink.LinkByName(linkname)
	if err == nil {
		// link exists
		return nil
	}

	l := netlink.NewLinkAttrs()
	for _, f := range fs {
		f(l)
	}
	l.Name = linkname

	mybridge := &netlink.Bridge{LinkAttrs: l}
	return netlink.LinkAdd(mybridge)
}

func DelLink(linkname string) error {
	ln, err := netlink.LinkByName(linkname)
	if err != nil {
		// link not  exists
		return nil
	}
	return netlink.LinkDel(ln)
}


func LinkUp(linkname string)error{
	l ,err:=netlink.LinkByName(linkname)
	if err != nil {
		return err
	}
	return netlink.LinkSetUp(l)
}

func LinkDown(linkname string)error{
	l ,err:=netlink.LinkByName(linkname)
	if err != nil {
		return err
	}
	return netlink.LinkSetDown(l)
}


func FindAllLink()([]netlink.Link,error){
	return netlink.LinkList()
}