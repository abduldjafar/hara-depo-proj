package transaksi

import (
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func ListTransaksi(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	transaksi := []mobile.TransaksiUang{}
	response := []mobile.TransaksiUang{}

	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	userJ := vars["kodeuser"]
	sort := vars["sort"]
	var qsort string

	dbOffset := page * 10

	if sort == "terbaru" {
		qsort = "create_date asc"
	} else if sort == "utang" {
		qsort = "utang desc"
	} else if sort == "total" {
		qsort = "total desc"
	} else {
		qsort = "nama_kasir asc"
	}

	if err := db.Where("kode_user=?", userJ).Offset(dbOffset).Limit(10).Order(qsort).Find(&transaksi).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, "error get datas")
		return
	}

	for _, data := range transaksi {
		var hutang mobile.Hutang
		if err := db.Where("id_hutang=?", data.IdHutang).First(&hutang).Error; err != nil {
			log.Println(err.Error())
			customResponse.RespondError(w, http.StatusInternalServerError, "error get datas")
			return
		}

		response = append(response, data)
	}
	customResponse.RespondJSON(w, 202, response)
}
