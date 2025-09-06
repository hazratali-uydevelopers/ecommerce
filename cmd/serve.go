package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	mux := http.NewServeMux()

	wrapedMux := manager.Wrap(
		mux, 
		middleware.Preflight, 
		middleware.Cors,
		middleware.Logger, 
	)

	initRoutes(mux, manager)

	fmt.Println("Server is runing at port :8080")
	err := http.ListenAndServe(":8080", wrapedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
