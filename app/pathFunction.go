package app

import "net/http"

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) GetQueryKategori(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET").Queries("page", "{page}")
}

func (app *App) GetQueryBarang(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET").Queries("page", "{page}", "kodeuser", "{kodeuser}")
}

func (app *App) GetBarangInformation(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET").Queries("idbarang", "{idbarang}", "kodeuser", "{kodeuser}")
}

func (app *App) GetQueryBarangSort(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET").Queries("page", "{page}", "kodeuser", "{kodeuser}", "sort", "{sort}")
}

func (app *App) DeleteBarang(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE").Queries("kodeuser", "{kodeuser}", "idbarang", "{idbarang}")
}

func (app *App) DeleteKategory(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE").Queries("kodeuser", "{kodeuser}", "kodekategory", "{kodekategory}")
}

func (app *App) UpdateBarang(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

func (app *App) GetQueryStruk(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET").Queries("idtransaksi", "{idtransaksi}", "kodeuser", "{kodeuser}")
}

func (app *App) GetQueryPelanggan(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET").Queries("idpelanggan", "{idpelanggan}", "kodeuser", "{kodeuser}")
}

func (app *App) GetQueryListPelanggan(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET").Queries("page", "{page}", "kodeuser", "{kodeuser}", "filter", "{filter}", "sort", "{sort}")
}

func (app *App) DeletePelanggan(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE").Queries("kodeuser", "{kodeuser}", "idpelanggan", "{idpelanggan}")
}

func (app *App) GetQueryNominal(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET").Queries("nominal", "{nominal}")
}
