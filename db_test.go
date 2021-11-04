package goshorturl

import (
	"database/sql"
	"testing"
)

func TestSqliteOperation(t *testing.T) {
	InitSQLite(true)

	_, err := CreateShortUrlTable()
	if err != nil {
		t.Error(err)
		return
	}

	u := ShortUrl{
		Id:           123456,
		ShortUrlCode: "ashiie",
		OriginalUrl:  "https://abc.ed/aas",
	}

	_, err = InsertShortUrl(u)
	if err != nil {
		t.Error(err)
		return
	}

	u2 := ShortUrl{
		Id:           56789,
		ShortUrlCode: "eurocse",
		OriginalUrl:  "https://dbc.es/tuy",
	}
	_, err = InsertShortUrl(u2)
	if err != nil {
		t.Error(err)
		return
	}

	us, err := SelectAllShortUrl()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(us)

	su, err := SelectById(56789)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(su)

	_, err = SelectById(777)
	if err != sql.ErrNoRows {
		t.Error(err)
		return
	}

	su, err = SelectByOriginalUrl("https://dbc.es/tuy")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(su)

	_, err = SelectByOriginalUrl("https://dbc.es/tuyyyy")
	if err != sql.ErrNoRows {
		t.Error(err)
		return
	}

	su, err = SelectByShortUrlCode("eurocse")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(su)

	_, err = SelectByShortUrlCode("eurocse1")
	if err != sql.ErrNoRows {
		t.Error(err)
		return
	}
}
