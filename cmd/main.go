package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/peterest/go-basic-ecom/cmd/api"
	"github.com/peterest/go-basic-ecom/config"
	"github.com/peterest/go-basic-ecom/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBHost,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(
		fmt.Sprintf(":%s", config.Env.Port),
		db,
	)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
