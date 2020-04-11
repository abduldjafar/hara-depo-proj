package user

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model"
	"hara-depo-proj/util"
	"io/ioutil"
	"net/http"
)

func RegisterUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.UserOtlet{}
	toko := model.Toko{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &toko)
	err2 := json.Unmarshal(body, &user)
	_ = err2

	if err := db.Save(&user).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, "error save user data")
		return
	}

	if err := db.Save(&toko).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, "error save data toko")
		return
	}
	util.RespondOk(w, 200, "Outlet berhasil Didaftarkan")
}
