package penjualan

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/otp"
	"hara-depo-proj/util/customResponse"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func JualBarang(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	jualan := mobile.RequestJualan{}
	Tbarang := mobile.TransaksiBarang{}
	uang := mobile.TransaksiUang{}
	hutang := mobile.Hutang{}

	body, err1 := ioutil.ReadAll(r.Body)

	err1 = json.Unmarshal(body, &jualan)

	if err1 != nil {
		log.Println(err1.Error())
	}

	kodeTransaksi := otp.RandomStringOTP(6)

	for _, data := range jualan.Barang {

		var stok mobile.Stok
		var stokBaru mobile.Stok
		var barang mobile.BarangOtlet

		Tbarang.IdBarang = data.IdBarang
		Tbarang.Idpelanggan = jualan.IdPelanggan
		Tbarang.IdTransaksi = kodeTransaksi
		Tbarang.Qty = data.Jumlah
		Tbarang.KodeUser = jualan.KodeUser
		if err := db.Save(&Tbarang).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := db.Where("id_barang=?", data.IdBarang).Find(&barang).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := db.Where("id_kategori=? AND kode_user=? AND id_barang=?",
			strconv.Itoa(barang.IdKategori), jualan.KodeUser, strconv.Itoa(data.IdBarang)).Order("time_created asc").First(&stok).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		stokBaru = stok
		stokBaru.IdStok = otp.RandomString(5)
		stokBaru.Penjualan = data.Jumlah
		stokBaru.StokAkhir = -data.Jumlah

		if err := db.Save(&stokBaru).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	uang.IdTransaksi = kodeTransaksi
	uang.IdPelanggan = jualan.IdPelanggan
	uang.DiskonNominal = jualan.DiskonNominal
	uang.DiskonDecimal = jualan.DiskonDecimal
	uang.PajakNominal = jualan.PPNNominal
	uang.PajakDecimal = jualan.PPNDecimal
	uang.Subtotal = jualan.Subtotal
	uang.Total = jualan.Total

	uang.NamaKasir = jualan.Struk.NamaKasir
	uang.JenisTransaksi = jualan.JenisTransaksi
	uang.TipeTransaksi = jualan.JenisPembayaran
	uang.Pembulatan = jualan.Pembulatan
	uang.KodeUser = jualan.KodeUser
	jualan.Struk.No = otp.RandomStringOTP(6)
	uang.KodeStruk = jualan.Struk.No
	tanggalPelunasan, _ := time.Parse("2006-01-02", jualan.TanggalPelunasan)
	uang.TanggalPelunasan = tanggalPelunasan
	uang.Note = jualan.Note
	uang.UangTunai = jualan.UangTunai

	if err := db.Save(&uang).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	} else {
		log.Println("sukses save data transaksi ==> kode " + kodeTransaksi)
	}
	if strings.ToLower(uang.TipeTransaksi) == "utang" {
		hutang.IdTransaksi = uang.IdTransaksi
		hutang.JatuhTempo = uang.TanggalPelunasan
		hutang.Pembayaran = uang.UangTunai
		hutang.SisaHutang = uang.Total - uang.UangTunai
		hutang.Status = "Utang"

		if err := db.Save(&hutang).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
			log.Println(err.Error())
			return
		} else {
			log.Println("sukses save data hutang  ==> kode " + kodeTransaksi)
		}

		if err := db.Model(&uang).Updates(mobile.TransaksiUang{IdHutang: hutang.IdHutang}).Error; err != nil {
			customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
			log.Println(err.Error())
			return
		}
	}

	jualan.Struk.Tanggal = uang.CreateDate.String()
	customResponse.RespondJSON(w, 200, jualan)

}
