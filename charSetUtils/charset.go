package charSetUtils

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/BurntSushi/toml"
)

type CHATSETTYPE int
type MARSHALTYPE int

const (
	BASE64_URL CHATSETTYPE = iota + 1
	BASE64_STD
	BASE64_RAW
	HEX

	JSON MARSHALTYPE = iota + 1
	XML
	TOML
)

var (
	MARSHALERR = fmt.Errorf("not found this MARSHALTYPE")
	CHARSETERR = fmt.Errorf("not found this CHATSETTYPE")
)

func Marshal(t MARSHALTYPE, data interface{}) ([]byte, error) {
	var (
		result []byte
		err    = MARSHALERR
	)
	switch t {
	case JSON:
		result, err = json.Marshal(data)
	case XML:
		result, err = xml.Marshal(data)
	case TOML:
		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(data)
		result = buf.Bytes()
	}
	return result, err
}
func UnMarshal(t MARSHALTYPE, endata []byte, data interface{}) error {
	var err error
	switch t {
	case JSON:
		err = json.Unmarshal(endata, data)
	case XML:
		err = xml.Unmarshal(endata, data)
	case TOML:
		_, err = toml.Decode(string(endata), data)
	default:
		err = fmt.Errorf("not found this %v MARSHALTYPE", t)
	}
	return err
}

func CharSetEncode(t CHATSETTYPE, data interface{}) (string, error) {
	var (
		result string
		err    = CHARSETERR
	)
	switch t {
	case BASE64_URL:
		v, ok := data.([]byte)
		if ok {
			result = base64.URLEncoding.EncodeToString(v)
		}
	case BASE64_STD:
		v, ok := data.([]byte)
		if ok {
			result = base64.StdEncoding.EncodeToString(v)
		}
	case BASE64_RAW:
		v, ok := data.([]byte)
		if ok {
			result = base64.RawStdEncoding.EncodeToString(v)
		}
	case HEX:
		v, ok := data.([]byte)
		if ok {
			result = hex.EncodeToString(v)
		}
	}
	return result, err
}

func CharSetDecode(t CHATSETTYPE, data string) ([]byte, error) {
	var (
		result []byte
		err    = CHARSETERR
	)
	switch t {
	case BASE64_URL:
		result, err = base64.URLEncoding.DecodeString(data)
	case BASE64_STD:
		result, err = base64.StdEncoding.DecodeString(data)
	case BASE64_RAW:
		result, err = base64.RawStdEncoding.DecodeString(data)
	case HEX:
		result, err = hex.DecodeString(data)
	}
	return result, err
}
