package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/siongui/goshorturl"
)

func HandleUrl(url string) {
	if u, err := goshorturl.SelectByOriginalUrl(url); err == nil {
		// The short url code for input URL exists. Return the code.
		fmt.Println("short url code: ", u.ShortUrlCode)
	} else {
		// The short url code for input URL does not exist.

		// create new id that is not in database and generate short link
		// code from the id.
		id := goshorturl.GetRandomId()
		_, err := goshorturl.SelectById(id)
		for err != sql.ErrNoRows {
			id := goshorturl.GetRandomId()
			_, err = goshorturl.SelectById(id)
		}

		row := goshorturl.ShortUrl{
			Id:           strconv.FormatUint(id, 10),
			ShortUrlCode: goshorturl.GetShortUrlCodeFromId(id),
			OriginalUrl:  url,
		}

		// insert newly created short link into database
		_, err = goshorturl.InsertShortUrl(row)
		if err != nil {
			fmt.Println("!!! Fail to insert into database", row, err)
			return
		}
		fmt.Printf("Inserted - Id: %s , Short Url Code: %s , URL: %s\n", row.Id, row.ShortUrlCode, row.OriginalUrl)
	}
}

func HandleShortUrlCode(code string) {
	if u, err := goshorturl.SelectByShortUrlCode(code); err == nil {
		// The short url code exists in the database. return original
		// URL.
		fmt.Println("original URL: ", u.OriginalUrl)
	} else {
		fmt.Println("HTTP 404 not found")
	}
}

func StartConsole() (err error) {
	goshorturl.InitSQLite(false)

	_, err = goshorturl.CreateShortUrlTable()
	if err != nil {
		return
	}

	for {
		fmt.Print("Please enter short url code or URL (ctrl+c to quit): ")
		var input string
		fmt.Scanln(&input)

		s := strings.TrimSpace(input)

		if strings.HasPrefix(s, "http") {
			HandleUrl(s)
		} else {
			HandleShortUrlCode(s)
		}
	}

	return
}

func main() {
	err := StartConsole()
	if err != nil {
		panic(err)
	}
}
