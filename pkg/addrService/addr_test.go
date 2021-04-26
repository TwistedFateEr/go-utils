package addrService

import (
	"testing"
)

func TestName(t *testing.T) {
	rs,err:=DeviceAndAddr()
	if err != nil {
		t.Fatal(err)
	}

	for link, addrs := range rs {
		t.Log(link,addrs)
	}


}