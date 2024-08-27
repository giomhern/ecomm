package main

import (
	"database/sql"
	"log"

	"github.com/giomhern/ecomm/cmd/api"
	"github.com/giomhern/ecomm/config"
	"github.com/giomhern/ecomm/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("Error found creating db", err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if error := server.Run(); error != nil {
		log.Fatal("Error found", error)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB is connected successfully")
}
