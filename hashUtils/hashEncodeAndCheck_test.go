package hashUtils

import "testing"


func TestEncode(t *testing.T) {
	data := []byte("hello encode")


	for _, encodetype := range []ENCODETYPE{
		SHA1, SHA256, SHA512, MD5,
	} {
		endata, err := EncodeData(encodetype, data)
		if err != nil {
			t.Fatal(err)
		}
		if !CheckData(encodetype, endata, data) {
			t.Fatal("data not eq  ", encodetype, endata, data)
		}
	}
}