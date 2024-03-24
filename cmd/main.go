package main

import (
	"fmt"
	"log"

	"com.github.com/peterest/go-basic-ecom/cmd/api"
	"com.github.com/peterest/go-basic-ecom/config"
	"com.github.com/peterest/go-basic-ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	// This is the entry point of the application
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
