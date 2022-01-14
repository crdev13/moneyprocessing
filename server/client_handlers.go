package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/crdev13/moneyprocessing/components/clients"
	"github.com/crdev13/moneyprocessing/components/clients/dto/input"
	"github.com/gorilla/mux"
)

func extractClientIDFromRequest(r *http.Request) (uint32, error) {
	vars := mux.Vars(r)
	identification := vars["CLIENTID"]
	var newID uint32
	idConverted, err := strconv.ParseInt(identification, 10, 64)
	if err != nil {
		return newID, fmt.Errorf("Error, cannot extract clientID")
	}
	newID = uint32(idConverted)
	return newID, nil
}

func (app *App) createClient(w http.ResponseWriter, r *http.Request) {
	request, err := input.MakeCreateClientRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	task, err := clients.NewCreateClient(
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

func (app *App) getClient(w http.ResponseWriter, r *http.Request) {
	clientID, err := extractClientIDFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	task, err := clients.NewGetClientByID(app.clientsRepository, clientID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	client, err := task.Execute()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, client)
}
