package kategori

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model"
	"hara-depo-proj/otp"
	"hara-depo-proj/util"
	"io/ioutil"
	"net/http"
)

func AddKategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	kategory := model.Kategory{}
	stok := model.Stok{}

	body, err1 := ioutil.ReadAll(r.Body)

	err1 = json.Unmarshal(body, &kategory)
	err2 := json.Unmarshal(body, &stok)
	_ = err2
	_ = err1

	defer r.Body.Close()

	if err := db.Save(&kategory).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	stok.IdKategori = kategory.KodeKategory
	stok.IdStok = otp.RandomString(5)

	if err := db.Save(&stok).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, kategory)

}

// output add kategori kategory + id
