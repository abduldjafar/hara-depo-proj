package barang

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"net/http"
	"strconv"
)

func ListBarangOtletOwner(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	barangs := []mobile.BarangJoinStok{}
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	userJ := vars["kodeuser"]
	sort := vars["sort"]
	var qsort string

	if sort == "terbaru" {
		qsort = "barang_otlet.time_created desc"
	} else if sort == "hargajual" {
		qsort = "barang_otlet.harga_jual desc"
	} else if sort == "stok" {
		qsort = "stok.stok_akhir desc"
	} else {
		qsort = "nama_barang asc"
	}

	dbOffset := page * 10

	if err := db.Table("barang_otlet").Select("barang_otlet.kode_user,barang_otlet.id_barang,barang_otlet"+
		".nama_barang,kategory.nama_kategory,barang_otlet.harga_jual,sum(stok.stok_akhir) as Jumlah").
		Joins("inner join stok on stok.id_barang = barang_otlet.id_barang").
		Joins("inner join kategory on kategory.kode_kategory = barang_otlet.id_kategori").
		Group("barang_otlet.kode_user,barang_otlet.id_barang,barang_otlet.nama_barang,kategory.nama_kategory,barang_otlet.harga_jual").
		Where("barang_otlet.kode_user=?", userJ).
		Offset(dbOffset).
		Limit(10).
		Order(qsort).
		Find(&barangs).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
		fmt.Println(resp)
	}
	customResponse.RespondJSON(w, http.StatusOK, barangs)

}
