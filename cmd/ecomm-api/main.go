package main

import (
	"log"

	"github.com/jumayevgadam/ecomm_mysql/internal/connection"
)

func main() {
	db, err := connection.NewDatabase()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	log.Println("Successfully connected to mysqlDB")

}
