package goshorturl

import (
	"testing"
)

func TestGenerateRandomShortUrlCode(t *testing.T) {
	id1 := GetRandomId()
	t.Log("id: ", id1)
	shorturl1 := GetShortUrlCodeFromId(id1)
	t.Log("short url: ", shorturl1)

	id2 := GetRandomId()
	t.Log("id: ", id2)
	shorturl2 := GetShortUrlCodeFromId(id2)
	t.Log("short url: ", shorturl2)
}
