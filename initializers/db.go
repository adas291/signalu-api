package initializers

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

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
