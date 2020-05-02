package user

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"io/ioutil"
	"net/http"
)

func RegisterUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := mobile.UserOtlet{}
	toko := mobile.Toko{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &toko)
	err2 := json.Unmarshal(body, &user)
	_ = err2

	if err := db.Save(&user).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, "error save user data")
		return
	}

	if err := db.Save(&toko).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, "error save data toko")
		return
	}
	customResponse.RespondOk(w, 200, "Outlet berhasil Didaftarkan")
}
