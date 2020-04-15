package suplier

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"io/ioutil"
	"net/http"
)

func AddSuplier(db1 *gorm.DB, db2 *gorm.DB, w http.ResponseWriter, r *http.Request) {

	suplier := mobile.Suplier{}

	body, err1 := ioutil.ReadAll(r.Body)

	err1 = json.Unmarshal(body, &suplier)
	_ = err1

	defer r.Body.Close()

	if err := db1.Save(&suplier).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

}
