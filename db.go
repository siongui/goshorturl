package goshorturl

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

// ShortUrl represents the structure of short links.
type ShortUrl struct {
	Id           uint64
	ShortUrlCode string
	OriginalUrl  string
}

var db *bun.DB
var ctx = context.Background()

// CreateShortUrlTable creates table in the database to store short links.
func CreateShortUrlTable() (sql.Result, error) {
	return db.NewCreateTable().Model((*ShortUrl)(nil)).Exec(ctx)
}

// InitSQLite initialize in-memory database to store data. The verbose flag
// indicates whether to print all queries to stdout.
func InitSQLite(verbose bool) (err error) {
	// Open an in-memory SQLite database.
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		return
	}

	sqldb.SetMaxOpenConns(1)

	// Create a Bun db on top of it.
	db = bun.NewDB(sqldb, sqlitedialect.New())

	// If you are using an in-memory database, you need to configure *sql.DB
	// to NOT close active connections. Otherwise, the database is deleted
	// when the connection is closed.
	//sqldb.SetMaxIdleConns(1000)
	//sqldb.SetConnMaxLifetime(0)

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(verbose)))

	return
}

// SelectAllShortUrl reads all short links records from the database.
func SelectAllShortUrl() (us []ShortUrl, err error) {
	err = db.NewSelect().
		Model(&us).
		OrderExpr("id ASC").
		Scan(ctx)
	return
}

// InsertShortUrl inserts one record of short links in the database.
func InsertShortUrl(u ShortUrl) (result sql.Result, err error) {
	return db.NewInsert().Model(&u).Exec(ctx)
}

// SelectById selects the record by id in the database.
func SelectById(id uint64) (us ShortUrl, err error) {
	err = db.NewSelect().
		Model(&us).
		Where("id = ?", id).
		Limit(1).
		Scan(ctx)
	return
}

// SelectByOriginalUrl selects the record by original URL in the database.
func SelectByOriginalUrl(oriurl string) (us ShortUrl, err error) {
	err = db.NewSelect().
		Model(&us).
		Where("original_url = ?", oriurl).
		Limit(1).
		Scan(ctx)
	return
}
