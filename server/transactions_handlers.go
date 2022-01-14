package server

import (
	"net/http"

	"github.com/crdev13/moneyprocessing/components/transactions"
	"github.com/crdev13/moneyprocessing/components/transactions/dto/input"
)

func (app *App) getTransactions(w http.ResponseWriter, r *http.Request) {
	accountID, err := extractAccountIDFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	task, err := transactions.NewGetTransactionsByAccount(
		app.clientsRepository,
		app.transactionsRepository,
		accountID,
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	account, err := task.Execute()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, account)
}

func (app *App) deposit(w http.ResponseWriter, r *http.Request) {
	request, err := input.MakeDepositRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	task, err := transactions.NewDepositMoney(
		app.clientsRepository,
		app.transactionsRepository,
		request,
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = task.Execute()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"msj": "success"})
}

func (app *App) withdraw(w http.ResponseWriter, r *http.Request) {
	request, err := input.MakeWithdrawRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	task, err := transactions.NewWithdrawMoney(
		app.clientsRepository,
		app.transactionsRepository,
		request,
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = task.Execute()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"msj": "success"})
}

func (app *App) transfer(w http.ResponseWriter, r *http.Request) {
	request, err := input.MakeTransferRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	task, err := transactions.NewTransferMoney(
		app.clientsRepository,
		app.transactionsRepository,
		request,
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = task.Execute()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"msj": "success"})
}
