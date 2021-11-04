package goshorturl

import (
	"encoding/json"
)

// This file is nothing to do with this Go short url. To be deleted.

// "content.th.email.change.name" = "value"
// to
//{
//  "content": {
//    "th": {
//      "email": {
//        "change": {
//          "name": "value"
//        }
//      }
//    }
//  }
//}

type MyStruct struct {
	Content struct {
		Th struct {
			Email struct {
				Change struct {
					Name string
				} `json:"change"`
			} `json:"email"`
		} `json:"th"`
	} `json:"content"`
}

func valueToJson(v string) (s string, err error) {
	val := MyStruct{}
	val.Content.Th.Email.Change.Name = v

	b, err := json.Marshal(val)
	s = string(b)
	return
}
