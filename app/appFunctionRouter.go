package app

import (
	"hara-depo-proj/controller/Account/otp"
	"hara-depo-proj/controller/Account/ownerotlet/barang"
	"hara-depo-proj/controller/Account/ownerotlet/kategori"
	"hara-depo-proj/controller/Account/ownerotlet/pembelian"
	"hara-depo-proj/controller/Account/ownerotlet/user"
	"net/http"
)

func (app *App) createNewUser(w http.ResponseWriter, r *http.Request) {
	otp.OtpUser(app.TbUserHara, w, r)
}

func (app *App) AddBarang(w http.ResponseWriter, r *http.Request) {
	barang.AddBarang(app.TbBarang, app.TbStok, w, r)
}

func (app *App) ListBarang(w http.ResponseWriter, r *http.Request) {
	barang.ListBarangOtletOwner(app.TbBarang, w, r)
}

func (app *App) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user.RegisterUserV2(app.TbUserOtlet, w, r)
}

func (app *App) ConfirmOtp(w http.ResponseWriter, r *http.Request) {
	otp.OtpUserConfirm(app.TbUserHara, w, r)
}

func (app *App) ListKategori(w http.ResponseWriter, r *http.Request) {
	kategori.ListKategori(app.TbKategori, w, r)
}

func (app *App) AddKategori(w http.ResponseWriter, r *http.Request) {
	kategori.AddKategory(app.TbKategori, w, r)
}

func (app *App) ListPersediaanBarang(w http.ResponseWriter, r *http.Request) {
	barang.ListPersediaanOtletOwner(app.TbStok, w, r)
}

func (app *App) BarangInformations(w http.ResponseWriter, r *http.Request) {
	barang.GetBarangInfo(app.TbStok, w, r)
}

func (app *App) PembelianBarang(w http.ResponseWriter, r *http.Request) {
	pembelian.BeliBarang(app.TbBarang, w, r)
}

func (app *App) DeleteBarangI(w http.ResponseWriter, r *http.Request) {
	barang.DeleteBarangOtletOwner(app.TbBarang, w, r)
}

func (app *App) DeleteKategoryI(w http.ResponseWriter, r *http.Request) {
	kategori.DeleteKategoryOtletOwner(app.TbKategori, w, r)
}

func (app *App) UpdateBarangI(w http.ResponseWriter, r *http.Request) {
	barang.UpdateBarang(app.TbKategori, w, r)
}

func (app *App) UpdateKategoryI(w http.ResponseWriter, r *http.Request) {
	kategori.UpdateKategory(app.TbKategori, w, r)
}
