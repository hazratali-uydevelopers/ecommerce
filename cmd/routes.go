package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /route",
		manager.With(
			http.HandlerFunc(handlers.Test),
			middleware.Arekta,
		))
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
			middleware.Arekta,
		))
	mux.Handle(
		"GET /products/{productID}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
			middleware.Arekta,
		))
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreatProduct),
			middleware.Arekta,
		))
}
