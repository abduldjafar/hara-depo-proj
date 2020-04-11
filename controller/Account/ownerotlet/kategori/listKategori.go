package kategori

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model"
	"hara-depo-proj/util"
	"net/http"
	"strconv"
)

func ListKategori(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	user := vars["kodeuser"]
	fmt.Println(user)

	dbOffset := page * 10

	kategori := []model.KategoriJoinStok{}

	if err := db.Table("kategory").Select("kategory.kode_kategory,kategory.nama_kategory,stok.kode_user"+
		",sum(stok.stok_akhir) as jumlah").
		Joins("right join stok on stok.id_kategori = kategory.kode_kategory").
		Group("(kategory.kode_kategory,kategory.nama_kategory,stok.kode_user)").
		Where("stok.kode_user = ?", user).Offset(dbOffset).Limit(10).Find(&kategori).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
		fmt.Println(resp)
	}
	util.RespondJSON(w, http.StatusOK, kategori)
}
