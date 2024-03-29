package kategori

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"net/http"
)

func DeleteKategoryOtletOwner(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	db.BlockGlobalUpdate(true)
	kategory := mobile.Kategory{}
	response := mobile.ResponseUniversal{}
	barang := mobile.BarangOtlet{}

	vars := mux.Vars(r)
	kodeuser := vars["kodeuser"]
	idkategory := vars["kodekategory"]

	if err := db.Where("kode_kategory=? AND kode_user=?", idkategory, kodeuser).
		First(&kategory).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, "Data Sudah Dihapus")
	} else {
		if err := db.Model(&barang).Where("id_kategori=? AND kode_user=?", idkategory, kodeuser).
			Update("id_kategori", 1).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, "Internasl Service Error")
		}

		if err := db.Where("kode_kategory=? AND kode_user=?", idkategory, kodeuser).
			Delete(&kategory).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, "Internasl Service Error")
		}

		response.Message = "Barang Terhapus"
		response.Status = 200
		customResponse.RespondJSON(w, 200, response)
	}

}
