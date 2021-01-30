package todos_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUsername = "USERNAME"
	dbPassword = "PASSWORD"
	dbHost     = "HOST"
	dbSchema   = "SCHEMA"
)

var (
	Client   *sql.DB
	username = os.Getenv(dbUsername)
	password = os.Getenv(dbPassword)
	host     = os.Getenv(dbHost)
	schema   = os.Getenv(dbSchema)
)

func init() {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database ready to accept connections")
}
