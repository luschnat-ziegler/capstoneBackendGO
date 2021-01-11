package app

import (
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