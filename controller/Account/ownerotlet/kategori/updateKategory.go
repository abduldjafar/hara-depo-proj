package kategori

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"log"
	"net/http"
	"strconv"
)

func UpdateKategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	kategory := mobile.Kategory{}
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
		log.Println(err.Error())
		customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	customResponse.RespondJSON(w, http.StatusOK, dataskategory)

}
