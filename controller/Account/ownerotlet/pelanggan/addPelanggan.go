package pelanggan

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/otp"
	"hara-depo-proj/util"
	"net/http"
	"time"
)

func AddPelanggan(db1 *gorm.DB, w http.ResponseWriter, r *http.Request) {

	pelanggan := mobile.Pelanggan{}
	transaksi := mobile.TransaksiUang{}

	KodeUser := r.FormValue("KodeUser")
	NamaSuplier := r.FormValue("NamaPelanggan")
	Gender := r.FormValue("Gender")
	TanggalLahir := r.FormValue("TanggalLahir")
	NoHp := r.FormValue("Hp")
	Email := r.FormValue("Email")
	Alamat := r.FormValue("Alamat")

	pelanggan.Nama = NamaSuplier
	pelanggan.KodeUser = KodeUser
	pelanggan.NoTlp = NoHp
	pelanggan.Email = Email
	pelanggan.Alamat = Alamat
	pelanggan.TanggalLahir = TanggalLahir
	pelanggan.Gender = Gender

	defer r.Body.Close()

	if err := db1.Save(&pelanggan).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db1.Select("id_pelanggan").Where("id_pelanggan=? AND kode_user=?", pelanggan.IDPelanggan, pelanggan.KodeUser).
		First(&transaksi).Error; err != nil {
		kodeTransaksi := otp.RandomStringOTP(6)
		transaksi.IdPelanggan = pelanggan.IDPelanggan
		transaksi.IdTransaksi = kodeTransaksi
		transaksi.CreateDate = time.Now()
		if err := db1.Save(&transaksi).Error; err != nil {
			util.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	util.RespondJSON(w, 202, pelanggan)

}
