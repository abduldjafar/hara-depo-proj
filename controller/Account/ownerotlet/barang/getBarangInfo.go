package barang

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model"
	"hara-depo-proj/util"
	"net/http"
)

func GetBarangInfo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	barangs := model.BarangJoinStok{}
	vars := mux.Vars(r)
	idbarang := vars["idbarang"]
	fmt.Println(idbarang)
	kodeuser := vars["kodeuser"]
	fmt.Println(kodeuser)

	if err := db.Table("barang_otlet").Select("barang_otlet.kode_user,barang_otlet.id_barang,barang_otlet"+
		".nama_barang,kategory.nama_kategory,barang_otlet.harga_jual,stok.stok_akhir").
		Joins("inner join stok on stok.id_barang = barang_otlet.id_barang").
		Joins("inner join kategory on kategory.kode_kategory = barang_otlet.id_kategori").
		Where("barang_otlet.kode_user=? AND barang_otlet.id_barang=?", kodeuser, idbarang).Find(&barangs).Error; err != nil {
		//var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
		util.RespondError(w, http.StatusNoContent, "No Items")
	}

	fmt.Println(barangs)
	util.RespondJSON(w, http.StatusOK, barangs)
}
