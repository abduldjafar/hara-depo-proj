package mobile

import "time"

type Hutang struct {
	IdHutang    int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	IdTransaksi string
	Status      string
	SisaHutang  float32
	Pembayaran  float32
	JatuhTempo  time.Time
}
