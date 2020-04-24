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
