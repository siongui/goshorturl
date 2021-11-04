package goshorturl

import (
	"math/rand"
	"time"

	base58 "github.com/itchyny/base58-go"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var encoding = base58.FlickrEncoding

// GetRandomId generate a pseudo random number.
func GetRandomId() uint64 {
	return r.Uint64()
}

// GetShortUrlCodeFromId uses base58 encoding to convert number to short url
// code.
func GetShortUrlCodeFromId(id uint64) string {
	return string(encoding.EncodeUint64(id))
}
