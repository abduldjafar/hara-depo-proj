package pelanggan

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"log"
	"net/http"
	"strconv"
	"time"
)

func UpdatePelanggan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	listAtrPelanggan := []string{"IDPelanggan", "KodeUser", "Nama", "Hp", "Email", "Alamat", "Gender", "TanggalLahir"}
	datas := map[string]interface{}{}
	pelanggan := mobile.Pelanggan{}

	for _, data := range listAtrPelanggan {
		if r.FormValue(data) != "" {
			idpelanggan, err := strconv.Atoi(r.FormValue(data))
			if err != nil || data == "Hp" {
				datas[data] = r.FormValue(data)
			} else {
				log.Println(idpelanggan)
				datas[data] = idpelanggan
			}

			if data == "Hp" {
				datas["NoTlp"] = r.FormValue("Hp")
			}
		}
	}
	datas["TimeCreated"] = time.Now()
	if err := db.Model(&pelanggan).Updates(datas).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}

	customResponse.RespondJSON(w, http.StatusOK, pelanggan)

}
