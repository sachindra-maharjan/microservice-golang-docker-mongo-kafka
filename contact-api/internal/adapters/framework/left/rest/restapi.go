package rest

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sm/contact-api/internal/models"
	"gopkg.in/mgo.v2/bson"
)

func (resta Adapter) Status(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	genResponse(w, http.StatusOK, map[string]string{"server": name, "result": "success"})
}

func (resta Adapter) AllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := resta.api.AllContacts()
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	genResponse(w, http.StatusOK, contacts)
}

func (resta Adapter) FindByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contact, err := resta.api.FindByID(params["id"])
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid Contact ID.")
		return
	}
	genResponse(w, http.StatusOK, contact)
}

func (resta Adapter) CreateContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
	contact.ID = bson.NewObjectId()
	if err := resta.api.CreateContact(contact); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	genResponse(w, http.StatusCreated, contact)
}	

func (resta Adapter) UpdateContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
	if err := resta.api.UpdateContact(contact); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	genResponse(w, http.StatusCreated, contact)
}

func (resta Adapter) DeleteContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
	if err := resta.api.DeleteContact(contact); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	genResponse(w, http.StatusCreated, map[string]string{"result":"success"})
}

func genResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func responseWithError(w http.ResponseWriter, code int, message string) {
	genResponse(w, code, map[string]string{"error": message})
}
