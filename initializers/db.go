package initializers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var LOCAL_DB *sql.DB

func OpenConnection() {

	username := "stud"
	password := "vLXCDmSG6EpEnhXX"
	hostname := "seklys.ila.lt"
	port := "3306"
	database := "LDB"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, hostname, port, database)

	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	DB = conn
}

func OpenLocalConnection() {

	database := "data.db"

	conn, err := sql.Open("sqlite3", database)

	if err != nil {
		log.Fatal(err)
	}

	LOCAL_DB = conn
}

func CreateSchema() error {

	schema := `
	CREATE TABLE IF NOT EXISTS matavimai (
	value INTEGER NOT NULL,
	x INTEGER NOT NULL,
	y INTEGER NOT NULL,
	distance REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS stiprumai (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	measurement INTEGER NOT NULL,
	sensor TEXT NOT NULL,
	strength INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS naudotojai (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	mac TEXT NOT NULL,
	sensor TEXT NOT NULL,
	strength INTEGER NOT NULL
);
`

	_, err := LOCAL_DB.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func NeedSeed() bool {

	var count int
	err := LOCAL_DB.QueryRow("SELECT COUNT(*) FROM matavimai").Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	return count == 0
}
