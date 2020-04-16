package barang

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/otp"
	"hara-depo-proj/util"
	"net/http"
	"strconv"
)

func AddBarang(db1 *gorm.DB, db2 *gorm.DB, w http.ResponseWriter, r *http.Request) {

	stok := mobile.Stok{}
	barang := mobile.BarangOtlet{}
	respon := mobile.ResponAddBarang{}

	KodeUser := r.FormValue("KodeUser")
	NamaBarang := r.FormValue("NamaBarang")
	Barcode := r.FormValue("Barcode")
	HargaJual := r.FormValue("HargaJual")
	HargaBeli := r.FormValue("HargaBeli")
	StokAwal := r.FormValue("StokAwal")
	StokAlarm := r.FormValue("StokAlarm")
	Deskripsi := r.FormValue("Deskripsi")
	IdKategory := r.FormValue("IdKategori")

	hargabeli, _ := strconv.ParseFloat(HargaBeli, 32)
	hargajual, _ := strconv.ParseFloat(HargaJual, 32)
	stokawal, _ := strconv.Atoi(StokAwal)
	stokalarm, _ := strconv.Atoi(StokAlarm)
	idkategory, _ := strconv.Atoi(IdKategory)

	idStok := otp.RandomString(5)
	stok.IdStok = idStok
	barang.KodeUser = KodeUser
	barang.NamaBarang = NamaBarang
	barang.Barcode = Barcode
	barang.HargaBeli = float32(hargabeli)
	barang.HargaJual = float32(hargajual)
	barang.Deskripsi = Deskripsi
	barang.IdKategori = idkategory
	stok.StokAwal = stokawal
	stok.StokAlarm = stokalarm
	stok.StokAkhir = stok.StokAwal
	stok.KodeUser = KodeUser
	stok.IdKategori = idkategory

	if err := db1.Save(&barang).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	stok.IDBarang = barang.IdBarang
	idbarang := strconv.Itoa(barang.IdBarang)
	barang.PhotoBarang = util.UploadPhotoAws(r, "PhotoBarang", barang.KodeUser, "barang", idbarang)

	if err := db1.Where("id_barang=?", idbarang).Save(&barang).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db2.Save(&stok).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respon.Barang = barang
	respon.StokBarang = stok
	util.RespondJSON(w, http.StatusOK, respon)

}

// output barang + id lengkap
