package main

import "net/http"

func (app *Application) Route(mux *http.ServeMux) *http.Server {
	mux.HandleFunc("GET /api/v1/user/{param}", app.getUser)
	mux.HandleFunc("GET /api/v1/product/{param}", app.getProduct)
	mux.HandleFunc("GET /api/v1/products", app.getProducts)
	mux.HandleFunc("POST /api/v1/product", app.createProduct)
	mux.HandleFunc("POST /api/v1/login", app.loginUser)
	mux.HandleFunc("POST /api/v1/register", app.registerUser)
	srv := http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}
	return &srv
}
