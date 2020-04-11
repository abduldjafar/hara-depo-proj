package barang

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model"
	"hara-depo-proj/util"
	"net/http"
	"strconv"
)

type response struct {
	Barang interface{}
	Stok   interface{}
}

func UpdateBarang(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	stok := model.Stok{}
	stokquery := model.Stok{}
	barang := model.BarangOtlet{}
	response := response{}

	datasBarang := map[string]interface{}{}
	datasStok := map[string]interface{}{}
	listkeysBarang := []string{"KodeUser", "NamaBarang", "Barcode", "HargaJual", "HargaBeli", "Deskripsi", "IdKategori", "IdBarang"}
	listkeysStok := []string{"StokAwal", "StokAlarm", "StokAkhir"}

	for _, data := range listkeysBarang {
		if r.FormValue(data) != "" {
			if data == "IdKategori" || data == "IdBarang" {
				idkategory, err := strconv.Atoi(r.FormValue(data))
				if err != nil {
					fmt.Println("========================")
				} else {
					datasBarang[data] = idkategory
				}
			} else if data == "HargaBeli" || data == "HargaJual" {
				harga, _ := strconv.ParseFloat(r.FormValue(data), 32)
				datasBarang[data] = float32(harga)
			} else {
				datasBarang[data] = r.FormValue(data)
			}
		}
	}

	for _, data := range listkeysStok {
		if r.FormValue(data) != "" {
			stoks, _ := strconv.Atoi(r.FormValue(data))
			datasStok[data] = stoks
		}
	}

	idbarang := r.FormValue("idbarang")

	_, _, err := r.FormFile("PhotoBarang")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	} else {
		datasBarang["PhotoBarang"] = util.UploadPhoto(r, "PhotoBarang", r.FormValue("KodeUser"), "barang", idbarang)
	}

	if err := db.Model(&barang).Updates(datasBarang).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Where("kode_user=? AND id_barang=?", r.FormValue("KodeUser"), r.FormValue("IdBarang")).
		First(&stokquery).Order("time_created asc").Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	datasStok["IdStok"] = stokquery.IdStok

	if err := db.Model(&stok).Updates(datasStok).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Barang = datasBarang
	response.Stok = datasStok

	util.RespondJSON(w, 200, response)
}
