package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/crdev13/moneyprocessing/components/clients"
	"github.com/crdev13/moneyprocessing/components/clients/dto/input"
	"github.com/gorilla/mux"
)

func extractAccountIDFromRequest(r *http.Request) (uint32, error) {
	vars := mux.Vars(r)
	identification := vars["ACCOUNTID"]
	var newID uint32
	idConverted, err := strconv.ParseInt(identification, 10, 64)
	if err != nil {
		return newID, fmt.Errorf("Error, cannot extract accountID")
	}
	newID = uint32(idConverted)
	return newID, nil
}

func (app *App) createAccount(w http.ResponseWriter, r *http.Request) {
	request, err := input.MakeCreateAccountRequest(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	task, err := clients.NewCreateAccount(
		app.clientsRepository,
		request,
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	response, err := task.Execute()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, response)
}

func (app *App) getAccount(w http.ResponseWriter, r *http.Request) {
	accountID, err := extractAccountIDFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	task, err := clients.NewGetAccountByID(app.clientsRepository, accountID)
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
