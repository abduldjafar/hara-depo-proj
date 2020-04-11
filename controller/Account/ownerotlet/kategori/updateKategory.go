package kategori

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model"
	"hara-depo-proj/util"
	"net/http"
	"strconv"
)

func UpdateKategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	kategory := model.Kategory{}
	dataskategory := map[string]interface{}{}

	listKategory := []string{"KodeUser", "KodeKategory", "NamaKategory"}
	for _, data := range listKategory {
		if r.FormValue(data) != "" {
			idkategory, err := strconv.Atoi(r.FormValue(data))
			if err != nil {
				dataskategory[data] = r.FormValue(data)
			} else {
				dataskategory[data] = idkategory
			}
		}
	}
	if err := db.Model(&kategory).Updates(dataskategory).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondJSON(w, http.StatusOK, dataskategory)

}
