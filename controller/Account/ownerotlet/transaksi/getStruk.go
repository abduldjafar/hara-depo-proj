package transaksi

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"net/http"
)

func GetStrukTransaksi(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idtransaksi := vars["idtransaksi"]
	kodeuser := vars["kodeuser"]

	response := mobile.ResponseGetStruk{}
	barangTrans := []mobile.BarangTransaksi{}
	detil := mobile.DetailTransaksi{}

	if err := db.Table("transaksi_uang").Select("kode_struk,tipe_transaksi,create_date").
		Where("id_transaksi=? AND kode_user=?", idtransaksi, kodeuser).Find(&detil).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Table("transaksi_uang").Where("id_transaksi=? AND kode_user=?", idtransaksi, kodeuser).Find(&response).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Table("barang_otlet").Select("transaksi_barang.kode_user,barang_otlet.nama_barang,barang_otlet.harga_jual,"+
		"transaksi_barang.qty,barang_otlet.harga_jual*transaksi_barang.qty as harga").
		Joins("inner join transaksi_barang on barang_otlet.id_barang = transaksi_barang.id_barang").
		Where("transaksi_barang.kode_user=? AND transaksi_barang.id_transaksi=?", kodeuser, idtransaksi).
		Find(&barangTrans).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Barang = barangTrans
	response.Detail = detil

	util.RespondJSON(w, 202, response)

}
