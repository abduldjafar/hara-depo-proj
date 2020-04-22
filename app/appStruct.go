package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router            *mux.Router
	SubRouter         *mux.Router
	TbUserHara        *gorm.DB
	TbBarang          *gorm.DB
	TbUserOtlet       *gorm.DB
	TbTokoOtlet       *gorm.DB
	TbKategori        *gorm.DB
	TbStok            *gorm.DB
	TbPelanggan       *gorm.DB
	TbSuplier         *gorm.DB
	TbTransaksiUang   *gorm.DB
	TbTransaksiBarang *gorm.DB
	Db                *gorm.DB
}
