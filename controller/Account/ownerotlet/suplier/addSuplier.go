package suplier

import (
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"net/http"

	"github.com/jinzhu/gorm"
)

func AddSuplier(db1 *gorm.DB, w http.ResponseWriter, r *http.Request) {

	suplier := mobile.Suplier{}

	KodeUser := r.FormValue("KodeUser")
	NamaSuplier := r.FormValue("NamaSuplier")
	NoHp := r.FormValue("Hp")
	Email := r.FormValue("Email")
	Alamat := r.FormValue("Alamat")

	suplier.Nama = NamaSuplier
	suplier.KodeUser = KodeUser
	suplier.NoTlp = NoHp
	suplier.Email = Email
	suplier.Alamat = Alamat

	defer r.Body.Close()

	if err := db1.Save(&suplier).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, 202, suplier)

}
