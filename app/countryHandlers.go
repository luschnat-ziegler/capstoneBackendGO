package app

import (
	"encoding/json"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"net/http"
)

type CountryHandlers struct {
	service service.CountryService
}


func (ch *CountryHandlers) getAllCountries(w http.ResponseWriter, _ *http.Request) {

	countries, err := ch.service.GetAll()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, countries)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}