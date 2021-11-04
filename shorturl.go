package goshorturl

import (
	"math/rand"
	"time"

	base58 "github.com/itchyny/base58-go"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var encoding = base58.FlickrEncoding

func getRandomId() uint64 {
	return r.Uint64()
}

func getShortUrl(id uint64) string {
	return string(encoding.EncodeUint64(id))
}
