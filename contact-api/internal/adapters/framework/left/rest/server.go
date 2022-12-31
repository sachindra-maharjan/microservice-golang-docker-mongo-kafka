package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sm/contact-api/internal/ports"
)

type Adapter struct {
		api ports.APIPort
}

func NewAdapter(api ports.APIPort) (*Adapter) {
	return &Adapter{api: api}
}

func (resta Adapter) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/status", resta.Status) 

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}

}