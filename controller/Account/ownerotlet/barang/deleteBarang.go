package barang

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"net/http"
)

func DeleteBarangOtletOwner(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	barang := mobile.BarangOtlet{}
	response := mobile.ResponseUniversal{}
	vars := mux.Vars(r)
	kodeuser := vars["kodeuser"]
	idbarang := vars["idbarang"]

	if err := db.Where("id_barang=? AND kode_user=?", idbarang, kodeuser).First(&barang).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, "Data Sudah Dihapus")
	} else {
		if err := db.Where("id_barang=? AND kode_user=?", idbarang, kodeuser).Delete(&barang).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, "Internasl Service Error")
		}

		response.Message = "Barang Terhapus"
		response.Status = 200
		customResponse.RespondJSON(w, 200, response)
	}
}
