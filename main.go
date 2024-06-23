package main

import (
	"itmx-test/config"
	db "itmx-test/data/database"
	"itmx-test/domain/customers"
	router "itmx-test/router/http"
	"net/http"

	customerStore "itmx-test/data/customers"
)

func main() {
	// get configuration stucts via .env file
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// establish DB connection
	db, err := db.ConnectSqlite(configuration.Database)
	if err != nil {
		panic(err)
	}
	// initialize repos and services using DI

	customerRepo := customerStore.New(db)
	customerSvc := customers.NewService(customerRepo)

	httpRouter := router.NewHTTPHandler(customerSvc)
	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
