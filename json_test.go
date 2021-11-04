package goshorturl

import (
	"testing"
)

// This file is nothing to do with this Go short url. To be deleted.

func TestValueToJson(t *testing.T) {
	jstring, err := valueToJson("value")
	if err != nil {
		t.Error(err)
	}
	t.Log(jstring)
}
