package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	clientsrepository "github.com/crdev13/moneyprocessing/components/clients/repository"
	transactionsrepository "github.com/crdev13/moneyprocessing/components/transactions/repository"
	"github.com/gorilla/handlers"

	"github.com/crdev13/moneyprocessing/server"
)

func main() {
	clients := clientsrepository.NewInMemoryClientsRepository()
	transactions := transactionsrepository.NewInMemoryTransactionsRepository()

	s, err := server.New(clients, transactions)
	if err != nil {
		fmt.Println(err)
		return
	}
	host := os.Getenv("HOST")
	fmt.Printf("The money processing server is on tap now: %v:8080\n", host)
	address := fmt.Sprintf("0.0.0.0:%v", 8080)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Origin", "Accept"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(address,
		handlers.CORS(
			allowedHeaders,
			allowedMethods,
			allowedOrigins,
		)(s.Router())))

}
