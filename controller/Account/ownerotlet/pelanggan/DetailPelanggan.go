package pelanggan

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"log"
	"net/http"
	"strconv"
)

func GetDetail(db *gorm.DB, w http.ResponseWriter, idpelanggan int, kodeuser string) {
	var errRespon error

	var response mobile.ResponseDetailPelanggan
	var PelangganSejak mobile.PlgganSejak
	var KunjunganTerakhir mobile.KungjnTrkhr

	if err := db.Table("pelanggan").Select("pelanggan.id_pelanggan,pelanggan.nama as Nama ,"+
		"pelanggan.no_tlp as Hp,sum(hutang.sisa_hutang) as Piutang,count(transaksi_uang.id_transaksi) as jumlah_transaksi,"+
		"sum(transaksi_uang.total) as nilai_transaksi,avg(transaksi_uang.total) rata_transaksi").
		Joins("left join transaksi_uang on pelanggan.id_pelanggan = transaksi_uang.id_pelanggan").
		Joins("left join hutang on transaksi_uang.id_transaksi = hutang.id_transaksi").
		Group("(pelanggan.id_pelanggan,Nama,Hp)").
		Where("pelanggan.id_pelanggan=? AND pelanggan.kode_user=?", idpelanggan, kodeuser).First(&response).Error; err != nil {
		log.Println(err.Error())
	}

	if err := db.Table("pelanggan").Select("time_created as pelanggan_sejak").Where("pelanggan.id_pelanggan=? "+
		"AND pelanggan.kode_user=?", idpelanggan, kodeuser).First(&PelangganSejak).Error; err != nil {
		log.Println(err.Error())
	}

	fmt.Println(idpelanggan)
	if errRespon = db.Table("transaksi_uang").Select("create_date as kunjungan_terakhir").Where("transaksi_uang.id_pelanggan=? "+
		"AND transaksi_uang.kode_user=?", idpelanggan, kodeuser).
		Order("kunjungan_terakhir desc").
		First(&KunjunganTerakhir).Error; errRespon != nil {
		log.Println(errRespon.Error())

	}

	if errRespon != nil {
		customResponse.RespondError(w, 500, errRespon.Error())
	} else {
		response.PelangganSejak = PelangganSejak.PelangganSejak
		response.JumlahTransaksi = response.JumlahTransaksi - 1
		response.KunjunganTerakhir = KunjunganTerakhir.KunjunganTerakhir
		customResponse.RespondJSON(w, 202, response)

	}
}
func DetailPelanggan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idpelanggan, _ := strconv.Atoi(vars["idpelanggan"])
	kodeuser := vars["kodeuser"]

	GetDetail(db, w, idpelanggan, kodeuser)

}
