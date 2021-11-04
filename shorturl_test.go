package goshorturl

import (
	"testing"
)

func TestGenerateRandomShortUrl(t *testing.T) {
	id1 := getRandomId()
	t.Log("id: ", id1)
	shorturl1 := getShortUrl(id1)
	t.Log("short url: ", shorturl1)

	id2 := getRandomId()
	t.Log("id: ", id2)
	shorturl2 := getShortUrl(id2)
	t.Log("short url: ", shorturl2)
}
