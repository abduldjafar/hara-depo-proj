package penjualan

import (
	"github.com/gorilla/mux"
	"hara-depo-proj/util/customResponse"
	"hara-depo-proj/util/rekomendasiUang"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Nominal []float64
}

func RekomendasiNominal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var response Response
	angka, err := strconv.ParseFloat(vars["nominal"], 64)

	if err != nil {
		log.Println(err.Error())
		customResponse.RespondError(w, 500, err.Error())
	} else {
		response.Nominal = rekomendasiUang.Nominal(angka)
		customResponse.RespondJSON(w, 202, response)
	}

}
