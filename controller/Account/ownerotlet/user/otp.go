package user

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"net/http"
)

func OtpUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := mobile.Userhara{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		customResponse.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	println(user.Hp)
}
