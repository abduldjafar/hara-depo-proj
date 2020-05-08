package transaksi

import (
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func ListTransaksi(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	response := []mobile.ResponListTransaksi{}
	responseAkhir := []mobile.ResponAkhirTransaksi{}
	responses := []mobile.ResponAkhirTransaksi{}

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

	if err := db.Table("transaksi_uang").Select("transaksi_uang.*,hutang.status as status_hutang,sum(hutang.sisa_hutang) as total_piutang").
		Joins("left join hutang on hutang.id_hutang = transaksi_uang.id_hutang").Where("kode_user=?", userJ).
		Group("transaksi_uang.*,transaksi_uang.id_transaksi,status_hutang").
		Offset(dbOffset).Limit(10).Order(qsort).Find(&response).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Table("transaksi_uang").Select("DISTINCT to_char(create_date, 'YYYY-MM-DD') as tanggal_transaksi").
		Joins("left join hutang on hutang.id_hutang = transaksi_uang.id_hutang").Where("kode_user=?", userJ).
		Offset(dbOffset).Limit(10).Find(&responseAkhir).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, data := range responseAkhir {
		for _, data2 := range response {
			var tanggal = strings.Split(data2.CreateDate.String(), " ")[0]
			if data.TanggalTransaksi == tanggal {
				data.Transaksi = append(data.Transaksi, data2)
			}
		}
		responses = append(responses, data)
	}

	customResponse.RespondJSON(w, 202, responses)
}
