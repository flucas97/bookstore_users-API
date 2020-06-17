package users_db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/flucas97/go-trainning/exercises/context/log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_root     = "mysql_users_root"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	root     = os.Getenv(mysql_users_root)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		root,
		schema,
	)

	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err) // Don't start the application if connection we have problem
	}

	// check if database is OK
	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println(context.Background(), "database successfuly conected")

}
