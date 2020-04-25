package mobile

import "time"

type DetailTransaksi struct {
	KodeStruk     string
	TipeTransaksi string
	CreateDate    time.Time
}
type BarangTransaksi struct {
	NamaBarang string
	HargaJual  float32
	Harga      float32
	Qty        float32
}

type ResponseGetStruk struct {
	Detail     DetailTransaksi
	Barang     []BarangTransaksi
	Pajak      float32
	Diskon     float32
	Subtotal   float32
	Total      float32
	Pembulatan float32
	Utang      float32
}
type TransaksiBarang struct {
	IdTransaksi string
	IdBarang    int
	IdSuplier   int
	Idpelanggan int
	Qty         float32
	KodeUser    string
}

type TransaksiUang struct {
	IdTransaksi    string `gorm:"PRIMARY_KEY"`
	KodeUser       string
	IdPelanggan    int
	NamaKasir      string
	KodeStruk      string
	JenisTransaksi string
	TipeTransaksi  string
	PajakNominal   float32
	PajakDecimal   float32
	DiskonNominal  float32
	DiskonDecimal  float32
	Subtotal       float32
	Total          float32
	Pembulatan     float32
	Utang          float32
	CreateDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
