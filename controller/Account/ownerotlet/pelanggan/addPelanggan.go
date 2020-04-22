package pelanggan

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"net/http"
)

func AddPelanggan(db1 *gorm.DB, w http.ResponseWriter, r *http.Request) {

	pelanggan := mobile.Pelanggan{}

	KodeUser := r.FormValue("KodeUser")
	NamaSuplier := r.FormValue("NamaPelanggan")
	NoHp := r.FormValue("Hp")
	Email := r.FormValue("Email")
	Alamat := r.FormValue("Alamat")

	pelanggan.Nama = NamaSuplier
	pelanggan.KodeUser = KodeUser
	pelanggan.NoTlp = NoHp
	pelanggan.Email = Email
	pelanggan.Alamat = Alamat

	defer r.Body.Close()

	if err := db1.Save(&pelanggan).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, 202, pelanggan)

}
