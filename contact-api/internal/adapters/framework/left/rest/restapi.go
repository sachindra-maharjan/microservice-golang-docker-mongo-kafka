package rest

import (
	"encoding/json"
	"net/http"
	"os"
)

func (resta Adapter) Status(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	genResponse(w, http.StatusOK, map[string]string{"server": name, "result": "success"})
}

func (resta Adapter)AllContacts(w http.ResponseWriter, r *http.Request) {

}

func (resta Adapter) FindByID(w http.ResponseWriter, r *http.Request) {

}

func (resta Adapter) FindByEmail(w http.ResponseWriter, r *http.Request) {

}

func (resta Adapter) CreateContact(w http.ResponseWriter, r *http.Request) {

}

func (resta Adapter) UpdateContact(w http.ResponseWriter, r *http.Request) {

}

func (resta Adapter) DeleteContact(w http.ResponseWriter, r *http.Request) {

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

