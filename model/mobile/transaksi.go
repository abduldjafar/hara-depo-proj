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
	Pembayaran float32
	Utang      float32
}
type TransaksiBarang struct {
	IdTransaksi  string
	IdBarang     int
	IdSuplier    int
	Idpelanggan  int
	Qty          float32
	KodeUser     string
	StatusBarang string
}

type TransaksiUang struct {
	IdTransaksi      string `gorm:"PRIMARY_KEY"`
	IdHutang         int
	IdRefund         int
	KodeUser         string
	IdPelanggan      int
	NamaKasir        string
	KodeStruk        string
	JenisTransaksi   string
	TipeTransaksi    string
	PajakNominal     float32
	PajakDecimal     float32
	DiskonNominal    float32
	DiskonDecimal    float32
	Subtotal         float32
	Total            float32
	Pembulatan       float32
	UangTunai        float32
	Note             string
	TanggalPelunasan time.Time
	CreateDate       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ResponListTransaksi struct {
	IdTransaksi      string `gorm:"PRIMARY_KEY"`
	IdHutang         int
	IdRefund         int
	KodeUser         string
	IdPelanggan      int
	NamaKasir        string
	KodeStruk        string
	JenisTransaksi   string
	TipeTransaksi    string
	PajakNominal     float32
	PajakDecimal     float32
	DiskonNominal    float32
	DiskonDecimal    float32
	Subtotal         float32
	Total            float32
	Pembulatan       float32
	UangTunai        float32
	Note             string
	TanggalPelunasan time.Time
	CreateDate       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	StatusHutang     string
	StatusRefund     string
	TotalPiutang     float32
}

type ResponAkhirTransaksi struct {
	TanggalTransaksi string
	Transaksi        []ResponListTransaksi
}
