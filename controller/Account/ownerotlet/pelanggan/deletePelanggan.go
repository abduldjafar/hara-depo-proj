package pelanggan

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"net/http"
)

func DeletePelanggan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	pelanggan := mobile.Pelanggan{}
	vars := mux.Vars(r)
	kodeuser := vars["kodeuser"]
	idpelanggan := vars["idpelanggan"]
	response := mobile.ResponseUniversal{}

	if err := db.Where("id_pelanggan=? AND kode_user=?", idpelanggan, kodeuser).
		First(&pelanggan).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Data Sudah Dihapus")
	} else {
		if err := db.Where("id_pelanggan=? AND kode_user=?", idpelanggan, kodeuser).
			Delete(&pelanggan).Error; err != nil {
			util.RespondError(w, http.StatusInternalServerError, "Internasl Service Error")
		}
		response.Message = "Pelanggan Terhapus"
		response.Status = 200
		util.RespondJSON(w, 200, response)
	}
}
