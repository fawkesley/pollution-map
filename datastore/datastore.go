package datastore

import (
	"database/sql"
	"log"

	"github.com/fawkesley/pollution-printouts/addresspollution"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	if err = createTables(); err != nil {
		log.Fatal(err)
	}
}

// Foo does nothing
func Foo() {}

func createTables() error {
	q := `CREATE TABLE IF NOT EXISTS addresses (
             "uuid"    UUID NOT NULL PRIMARY KEY,        
             "address_line_1" TEXT,
             "address_line_2" TEXT,
             "address_postcode" TEXT,
             "pm2_5"   FLOAT,
             "pm10"    FLOAT,
             "no2"     FLOAT
     );`

	statement, err := db.Prepare(q)
	if err != nil {
		return err
	}
	_, err = statement.Exec() // Execute SQL Statements
	if err != nil {
		return err
	}
	return nil
}

// SaveAddress saves the pollution levels for the given address
func SaveAddress(addr addresspollution.Address, levels addresspollution.PollutionLevels) error {
	q := `INSERT INTO addresses(uuid, address, pm2_5, pm10, no2) VALUES(?, ?, ?, ?, ?)`

	statement, err := db.Prepare(q)
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		addr.ID,
		levels.FormattedAddress,
		levels.Pm2_5,
		levels.Pm10,
		levels.No2,
	)
	if err != nil {
		return err
	}
	return nil
}
