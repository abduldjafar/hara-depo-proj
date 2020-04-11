package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/config"
	"hara-depo-proj/model"
	"log"
	"net/http"
)

func (app *App) Initialize() {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	db, err := gorm.Open("postgres", "host="+baseConfig.Postgres.Url+" port="+baseConfig.Postgres.Port+""+
		" user="+baseConfig.Postgres.User+" dbname="+baseConfig.Postgres.Db+" password="+baseConfig.Postgres.Password+
		" sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	app.TbUserHara = model.DBMigrationAccount(db, &model.Userhara{})
	app.TbBarang = model.DBMigrationAccount(db, &model.BarangOtlet{})
	app.TbUserOtlet = model.DBMigrationAccount(db, model.UserOtlet{})
	app.TbTokoOtlet = model.DBMigrationAccount(db, model.Toko{})
	app.TbKategori = model.DBMigrationAccount(db, model.Kategory{})
	app.TbStok = model.DBMigrationAccount(db, model.Stok{})
	app.TbPelanggan = model.DBMigrationAccount(db, model.Pelanggan{})
	app.TbSuplier = model.DBMigrationAccount(db, model.Suplier{})
	app.Router = mux.NewRouter()
	//a.SubRouter = a.Router.PathPrefix("/auth").Subrouter()
	//a.SubRouter.Use(auth.JwtVerify)
	app.setRouters()
	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./dist/")))
	app.Router.PathPrefix("/swaggerui/").Handler(sh)
}
