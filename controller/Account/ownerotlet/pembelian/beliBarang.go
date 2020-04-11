package pembelian

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model"
	"hara-depo-proj/util"
	"io/ioutil"
	"net/http"
)

func getbarang(db *gorm.DB, idbarang string, barang model.BarangOtlet) model.BarangOtlet {
	if err := db.Where("barang_otlet.id_barang = ?", idbarang).Find(&barang).Error; err != nil {
		return barang
	}
	return barang
}

func BeliBarang(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	beli := []model.RequestPembelian{}
	rincian := model.RincianBelanja{}
	listBelanja := []model.Rincianbelanja{}
	belanja := model.Rincianbelanja{}
	barang := model.BarangOtlet{}
	barangTemp := model.BarangOtlet{}
	listharga := []float32{}

	body, err1 := ioutil.ReadAll(r.Body)

	err1 = json.Unmarshal(body, &beli)
	_ = err1

	defer r.Body.Close()
	fmt.Println(beli)
	for _, data := range beli {
		barang = getbarang(db, data.IdBarang, barangTemp)
		listharga = append(listharga, float32(data.Jumlah)*barang.HargaBeli)
		belanja.NamaBarang = barang.NamaBarang
		belanja.HargaSatuan = barang.HargaBeli
		belanja.HargaTotal = float32(data.Jumlah) * barang.HargaBeli
		belanja.Jumlah = data.Jumlah
		listBelanja = append(listBelanja, belanja)

	}

	var sum = float32(0.0)
	for _, num := range listharga {
		sum += num
	}

	rincian.Rincian = listBelanja
	rincian.TotalPembelanjaan = sum
	util.RespondJSON(w, http.StatusOK, rincian)
}
