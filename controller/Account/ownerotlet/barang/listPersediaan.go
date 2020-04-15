package barang

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"net/http"
	"strconv"
)

func ListPersediaanOtletOwner(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	barangs := []mobile.KategoryJStockBrng{}
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	userk := vars["kodeuser"]

	dbOffset := page * 10

	if err := db.Table("barang_otlet").Select("barang_otlet.kode_user,barang_otlet.id_barang,barang_otlet"+
		".nama_barang,kategory.nama_kategory,stok.penjualan,stok.stok_akhir").
		Joins("inner join stok on stok.id_barang = barang_otlet.id_barang").
		Joins("inner join kategory on barang_otlet.id_kategori = kategory.kode_kategory").
		Where("barang_otlet.kode_user= ?", userk).Offset(dbOffset).Limit(10).Find(&barangs).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
		fmt.Println(resp)
	}

	fmt.Println(barangs)
	util.RespondJSON(w, http.StatusOK, barangs)

}
