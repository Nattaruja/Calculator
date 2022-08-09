package main

import (
	"app/customer"
	"app/router"
	"fmt"
	"net/http"
)

func initDeps() {
	customer.InitDB()
}

func main() {
	initDeps()

	router := router.GenerateRouter()

	srv := &http.Server{
		Addr:    ":8090",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
