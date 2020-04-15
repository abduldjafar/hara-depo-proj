package kategori

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
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

	kategori := []mobile.KategoriJoinStokBr{}
	barangs := []mobile.BarangId{}
	response := []mobile.KategoriJoinStokBr{}

	if err := db.Table("kategory").Select("kategory.kode_kategory,kategory.nama_kategory,stok.kode_user"+
		",sum(stok.stok_akhir) as jumlah").
		Joins("right join stok on stok.id_kategori = kategory.kode_kategory").
		Group("(kategory.kode_kategory,kategory.nama_kategory,stok.kode_user)").
		Where("stok.kode_user = ?", user).Offset(dbOffset).Limit(10).Find(&kategori).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
		fmt.Println(resp)
		util.RespondError(w, http.StatusInternalServerError, err.Error())

	}

	for _, data := range kategori {
		if err := db.Table("barang_otlet").Select("barang_otlet.id_barang").Joins("inner join stok on stok.id_barang = barang_otlet.id_barang").
			Joins("inner join kategory on kategory.kode_kategory = barang_otlet.id_kategori").
			Where("barang_otlet.kode_user=? AND barang_otlet.id_kategori=?", data.KodeUser, data.KodeKategory).
			Offset(dbOffset).
			Limit(10).
			Order("nama_barang asc").
			Find(&barangs).Error; err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
			fmt.Println(resp)
			util.RespondError(w, http.StatusInternalServerError, err.Error())

		}

		data.Barang = barangs
		response = append(response, data)
	}
	util.RespondJSON(w, http.StatusOK, response)
}
