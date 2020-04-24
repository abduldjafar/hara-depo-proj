package pelanggan

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util"
	"net/http"
	"strconv"
)

func DetailPelanggan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idpelanggan, _ := strconv.Atoi(vars["idpelanggan"])
	kodeuser := vars["kodeuser"]

	response := mobile.ResponseDetailPelanggan{}
	PelangganSejak := mobile.PlgganSejak{}
	KunjunganTerakhir := mobile.KungjnTrkhr{}

	if err := db.Table("pelanggan").Select("pelanggan.id_pelanggan,pelanggan.nama as Nama ,"+
		"pelanggan.no_tlp as Hp,sum(transaksi_uang.utang) as Piutang,count(transaksi_uang.id_transaksi) as jumlah_transaksi,"+
		"sum(transaksi_uang.total) as nilai_transaksi,avg(transaksi_uang.total) rata_transaksi").
		Joins("left join transaksi_uang on pelanggan.id_pelanggan = transaksi_uang.id_pelanggan").
		Group("(pelanggan.id_pelanggan,Nama,Hp)").
		Where("pelanggan.id_pelanggan=? AND pelanggan.kode_user=?", idpelanggan, kodeuser).First(&response).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	if err := db.Table("pelanggan").Select("time_created as pelanggan_sejak").Where("pelanggan.id_pelanggan=? "+
		"AND pelanggan.kode_user=?", idpelanggan, kodeuser).First(&PelangganSejak).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	if err := db.Table("transaksi_uang").Select("create_date as kunjungan_terakhir").Where("transaksi_uang.id_pelanggan=? "+
		"AND transaksi_uang.kode_user=?", idpelanggan, kodeuser).
		Order("kunjungan_terakhir desc").
		First(&KunjunganTerakhir).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
	}
	response.PelangganSejak = PelangganSejak.PelangganSejak
	response.KunjunganTerakhir = KunjunganTerakhir.KunjunganTerakhir
	util.RespondJSON(w, 202, response)
}

/*
{
    "Nama":"",
    "Hp":"",
    "PelangganSejak":"",
    "KunjunganTerakhir":"",
    "Piutang":0,
    "JumlahTransaksi":"",
    "NilaiTransaksi":"",
    "RataTransaksi":

}

select pelanggan.id_pelanggan,pelanggan.nama as Nama ,pelanggan.no_tlp as Hp,sum(transaksi_uang.utang) as total_piutang,count(transaksi_uang.id_transaksi) as jumlah_transaksi,sum(transaksi_uang.total) as nilai_transaksi,avg(transaksi_uang.total) rata_transaksi
from pelanggan left join transaksi_uang on pelanggan.id_pelanggan = transaksi_uang.id_pelanggan where pelanggan.id_pelanggan=1
group by (pelanggan.id_pelanggan,Nama,Hp);

*/
