package barang

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"log"
	"net/http"
	"strconv"
)

type response struct {
	Barang interface{}
	Stok   interface{}
}

func UpdateBarang(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	stok := mobile.Stok{}
	stokquery := mobile.Stok{}
	barang := mobile.BarangOtlet{}
	barangResponse := mobile.BarangOtlet{}
	//stokResponse := model.Stok{}
	response := response{}

	datasBarang := map[string]interface{}{}
	datasStok := map[string]interface{}{}
	listkeysBarang := []string{"KodeUser", "NamaBarang", "Barcode", "HargaJual", "HargaBeli", "Deskripsi", "IdKategori", "IdBarang"}
	listkeysStok := []string{"StokAwal", "StokAlarm", "StokAkhir"}

	for _, data := range listkeysBarang {
		if r.FormValue(data) != "" {
			if data == "IdKategori" || data == "IdBarang" {
				id, err := strconv.Atoi(r.FormValue(data))
				if err != nil {
					log.Println(err.Error())
				} else {
					datasBarang[data] = id
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
		log.Println(err.Error())
	} else {
		datasBarang["PhotoBarang"] = util.UploadPhotoAws(r, "PhotoBarang", r.FormValue("KodeUser"), "barang", idbarang)
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

	if err := db.Where("kode_user=? AND id_barang=?", r.FormValue("KodeUser"), r.FormValue("IdBarang")).
		Find(&barangResponse).Error; err != nil {
		fmt.Println("ahay")
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return

	}
	response.Barang = barangResponse
	response.Stok = stok

	util.RespondJSON(w, 200, response)
}
