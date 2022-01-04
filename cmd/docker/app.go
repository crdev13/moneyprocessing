package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	clientsrepository "github.com/crdev13/moneyprocessing/components/clients/repository"
	transactionsrepository "github.com/crdev13/moneyprocessing/components/transactions/repository"
	"github.com/gorilla/handlers"

	"github.com/crdev13/moneyprocessing/server"
)

func newPostgreSQLDatabaseConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	databaseHost := os.Getenv("DBHOST")

	databaseName := os.Getenv("DBNAME")
	databaseUser := os.Getenv("DBUSER")
	databasePass := os.Getenv("DBPASS")

	dbConnURL := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", databaseUser, databasePass, databaseHost, databaseName)

	dbConn, err := newPostgreSQLDatabaseConnection(dbConnURL)
	if err != nil || dbConn == nil {
		fmt.Println(err)
		fmt.Println("Cannot connect to posgreSQL database")
		return
	}
	clients, err := clientsrepository.NewInPostgreSQLClientsRepository(dbConn)
	if err != nil {
		fmt.Println(err)
		return
	}
	transactions, err := transactionsrepository.NewInPostgreSQLTransactionsRepository(dbConn)
	if err != nil {
		fmt.Println(err)
		return
	}

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
