package hashUtils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

type ENCODETYPE int

const (
	SHA1 ENCODETYPE = iota + 1
	SHA256
	SHA512
	MD5
)

var (
	encoder = new(encode)
)

type encode struct{}

func (e *encode) getEncoder(encodeType ENCODETYPE) hash.Hash {
	var encoder hash.Hash

	switch encodeType {
	case SHA1:
		encoder = sha1.New()
	case SHA256:
		encoder = sha256.New()
	case SHA512:
		encoder = sha512.New()
	case MD5:
		encoder = md5.New()
	default:
		encoder = sha256.New()
	}
	return encoder
}

func (e *encode) Encode(encodeType ENCODETYPE, data []byte) (enData []byte, err error) {
	encoder := e.getEncoder(encodeType)
	_, err = encoder.Write(data)
	if err != nil {
		return nil, err
	}
	return encoder.Sum(nil), nil
}

func EncodeData(encodeType ENCODETYPE, data []byte) (enData []byte, err error) {
	return encoder.Encode(encodeType, data)
}

func CheckData(encodeType ENCODETYPE, enData []byte, data []byte) bool {
	endata, err := encoder.Encode(encodeType, data)
	if err != nil {
		return false
	}
	return bytes.Equal(endata, enData)
}
