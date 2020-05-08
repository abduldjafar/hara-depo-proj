package penjualan

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/controller/Account/ownerotlet/pelanggan"
	"hara-depo-proj/model/mobile"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

func UtangJualan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	hutang := []mobile.Hutang{}
	bayarHutang := mobile.RequestBayarHutang{}

	body, err1 := ioutil.ReadAll(r.Body)

	err1 = json.Unmarshal(body, &bayarHutang)

	if err1 != nil {
		log.Println(err1.Error())
	}
	var sisa float32
	uang := bayarHutang.Pembayaran

	sisa = uang

	if err := db.Table("pelanggan").Select("hutang.*").
		Joins("inner join transaksi_uang on transaksi_uang.id_pelanggan = pelanggan.id_pelanggan").
		Joins("inner join hutang  on transaksi_uang.id_transaksi = hutang.id_transaksi").
		Where("transaksi_uang.kode_user=? AND pelanggan.id_pelanggan=?", bayarHutang.KodeUser, bayarHutang.IdPelanggan).
		Order("transaksi_uang.create_date desc").
		Find(&hutang).Error; err != nil {
		log.Println(err.Error())
	}

	for sisa > 0 {
		for _, dataHutang := range hutang {
			fmt.Println(sisa)
			if dataHutang.SisaHutang > 0 {
				var bayarngHutang mobile.Hutang
				dataHutang.SisaHutang = float32(math.Abs(float64(uang - dataHutang.SisaHutang)))
				sisa = dataHutang.SisaHutang
				bayarngHutang = dataHutang
				if err := db.Save(&bayarngHutang).Error; err != nil {
					log.Println(err.Error())
				}
			} else {
				sisa = 0
			}
			if sisa == 0 {
				break
			}
		}

	}

	pelanggan.GetDetail(db, w, bayarHutang.IdPelanggan, bayarHutang.KodeUser)
}
