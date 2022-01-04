package server

import (
	"net/http"

	clientsrepository "github.com/crdev13/moneyprocessing/components/clients/repository"
	transactionsrepository "github.com/crdev13/moneyprocessing/components/transactions/repository"
	"github.com/gorilla/mux"
)

type Server interface {
	Router() http.Handler
}

type App struct {
	router                 http.Handler
	clientsRepository      clientsrepository.ClientsRepository
	transactionsRepository transactionsrepository.TransactionsRepository
}

func New(
	clients clientsrepository.ClientsRepository,
	transactions transactionsrepository.TransactionsRepository,
) (Server, error) {
	r := mux.NewRouter()
	a := &App{
		clientsRepository:      clients,
		transactionsRepository: transactions,
	}
	a.loadHandlers(r)
	a.router = r

	return a, nil
}
func (a *App) Router() http.Handler {
	return a.router
}

func (a *App) loadHandlers(router *mux.Router) {
	router.HandleFunc("/api/clients", a.createClient).Methods(http.MethodPost)
	router.HandleFunc("/api/clients/id/{CLIENTID:[1-9][0-9]{0,8}}", a.getClient).Methods(http.MethodGet)
	router.HandleFunc("/api/accounts", a.createAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/accounts/id/{ACCOUNTID:[1-9][0-9]{0,8}}", a.getAccount).Methods(http.MethodGet)
	router.HandleFunc("/api/transactions/id/{ACCOUNTID:[1-9][0-9]{0,8}}", a.getTransactions).Methods(http.MethodGet)
	router.HandleFunc("/api/transactions/deposit", a.deposit).Methods(http.MethodPost)
	router.HandleFunc("/api/transactions/withdraw", a.withdraw).Methods(http.MethodPost)
	router.HandleFunc("/api/transactions/transfer", a.transfer).Methods(http.MethodPost)
}
