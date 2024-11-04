package main

import (
	"log"

	"github.com/ianschenck/envflag"
	"github.com/jumayevgadam/ecomm_mysql/internal/connection"
	"github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/handler"
	"github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/server"
	"github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/storer"
)

const minSecretKeySize = 32

func main() {
	var secretKey = envflag.String("SECRET_KEY", "012345678901234567890123456789012345678901", "for-ecommerce")
	envflag.Parse()

	if len(*secretKey) < minSecretKeySize {
		log.Fatalf("SECRET_KEY must be at least %d characters", minSecretKeySize)
	}

	db, err := connection.NewDatabase()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	log.Println("Successfully connected to mysqlDB")

	// do something with the database
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv, *secretKey)
	handler.RegisterRoutes(hdl)
	if err := handler.Start(":1323"); err != nil {
		log.Fatalf("error starting port: %v", err.Error())
	}
}
