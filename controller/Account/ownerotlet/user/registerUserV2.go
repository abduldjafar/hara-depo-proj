package user

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"hara-depo-proj/util/image"
	"net/http"
	"sync"
)

func RegisterUserV2(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var State = &mobile.State{&sync.Mutex{}, map[string]string{}}

	State.Lock()
	defer State.Unlock()

	KodeUser := r.FormValue("KodeUser")
	NamaLengkap := r.FormValue("NamaLengkap")
	NomorKtp := r.FormValue("NomorKtp")
	TanggalLahir := r.FormValue("TanggalLahir")
	JenisKelamin := r.FormValue("JenisKelamin")
	Email := r.FormValue("Email")
	NamaToko := r.FormValue("NamaToko")
	AlamatToko := r.FormValue("AlamatToko")
	KodeReferensi := r.FormValue("KodeReferensi")
	Hp := r.FormValue("Hp")

	user := mobile.UserOtlet{}
	toko := mobile.Toko{}
	response := mobile.ResponseRegister{}

	user.KodeUser = KodeUser
	user.NamaLengkap = NamaLengkap
	user.NomorKtp = NomorKtp
	user.TanggalLahir = TanggalLahir
	user.JenisKelamin = JenisKelamin
	user.FotoKtp = image.UploadPhotoAws(r, "FotoKtp", KodeUser, "user", "ktp")
	user.SelfieKtp = image.UploadPhotoAws(r, "SelfieKtp", KodeUser, "user", "ktp")
	user.FotoKontrak = image.UploadPhotoAws(r, "FotoKontrak", KodeUser, "user", "kontrak")
	user.Email = Email
	user.Hp = Hp
	toko.KodeUser = KodeUser
	toko.NamaToko = NamaToko
	toko.FotoToko = image.UploadPhotoAws(r, "FotoToko", KodeUser, "user", "Toko")
	user.FotoToko = image.UploadPhotoAws(r, "FotoToko", KodeUser, "user", "Toko")
	toko.AlamatToko = AlamatToko
	toko.KodeReferensi = KodeReferensi
	user.Status = "Disabled"

	if err := db.Save(&user).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, "error save user data")
		return
	}

	if err := db.Save(&toko).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, "error save data toko")
		return
	}
	response.User = user
	response.TokoUser = toko
	customResponse.RespondJSON(w, 200, response)

}
