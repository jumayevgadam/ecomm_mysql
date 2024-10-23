package main

import (
	"context"
	"log"

	"github.com/jumayevgadam/ecomm_mysql/internal/connection"
	"github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/handler"
	"github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/server"
	"github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/storer"
)

func main() {
	db, err := connection.NewDatabase()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	log.Println("Successfully connected to mysqlDB")

	// do something with the database
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(context.Background(), srv)
	handler.RegisterRoutes(hdl)
	handler.Start(":8000")
}
