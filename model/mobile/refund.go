package mobile

import "time"

type Refund struct {
	IdRefund     int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	IdTransaksi  int
	JumlahRefund float32
	Date         time.Time
}
