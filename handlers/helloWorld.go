package handlers

import (
	"fmt"
	"net/http"

	"github.com/wkrause13/go-transactionid-tutorial/middleware"
	"github.com/wkrause13/go-transactionid-tutorial/repositories"
)

type HelloWorldHandler struct {
	dataRepo repositories.GenericDataRepo
}

func NewHelloWorldHandler(repo repositories.GenericDataRepo) HelloWorldHandler {
	return HelloWorldHandler{dataRepo: repo}
}

func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	transID := r.Context().Value(middleware.TransactionId).(int)
	repo := h.dataRepo.CloneWithTransID(transID)
	users := repo.ListUsers()
	fmt.Fprintf(w, "Hello World: "+users[0])
}
