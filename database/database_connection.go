package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

const (
	server   = "DESKTOP-7UTUNVD"
	port     = 1433
	database = "ProjectSwagger"
)

var DB *sql.DB

func DatabaseInit() {
	connectionString := fmt.Sprintf("server=%s;port=%d;database=%s", server, port, database)
	var err error
	DB, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to SQL Server")
}
