package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/wkrause13/go-transactionid-tutorial/handlers"
	"github.com/wkrause13/go-transactionid-tutorial/middleware"
	"github.com/wkrause13/go-transactionid-tutorial/repositories"
)

const (
	GET = "GET"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.TransactionIdMiddleware)
	dataRepo := repositories.NewDataRepo()

	handler := handlers.NewHelloWorldHandler(dataRepo)

	r.Method(GET, "/", handler)
	http.ListenAndServe(":3000", r)

}
