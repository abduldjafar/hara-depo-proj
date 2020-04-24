package pelanggan

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"net/http"
	"strconv"
)

func ListPelanggan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	kodeuser := vars["kodeuser"]

	listPenjualan := []mobile.ListPenjualan{}
	response := mobile.ResponseListPenjualan{}
	jumlahPiutang := mobile.TotalUtang{}

	dbOffset := page * 10

	if err := db.Table("pelanggan").Select(" pelanggan.nama,pelanggan.id_pelanggan,sum(utang) as utang").
		Joins("left join transaksi_uang on pelanggan.id_pelanggan = transaksi_uang.id_pelanggan").
		Group("pelanggan.nama,pelanggan.id_pelanggan").
		Where("pelanggan.kode_user=?", kodeuser).
		Offset(dbOffset).
		Limit(10).
		Find(&listPenjualan).Error; err != nil {
		//var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
		util.RespondError(w, http.StatusInternalServerError, err.Error())

	}

	if err := db.Table("transaksi_uang").Select("sum(utang) as total_piutang").
		Where("transaksi_uang.kode_user=?", kodeuser).
		First(&jumlahPiutang).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	response.Penjualan = listPenjualan
	response.TotalPiutang = jumlahPiutang

	util.RespondJSON(w, 202, response)
}
