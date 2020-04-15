package mobile

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
	"time"
)

// structur of Userhara table
type Userhara struct {
	Id     int
	Nama   string
	Hp     string `gorm:"unique;not null"`
	Role   string
	Status string
}

type BarangOtlet struct {
	KodeUser    string
	IdBarang    int `gorm:"primary_key;AUTO_INCREMENT"`
	NamaBarang  string
	IdPelanggan string
	IdSuplier   string
	IdKategori  int // ganti ID kategory
	Barcode     string
	HargaJual   float32
	HargaBeli   float32
	Deskripsi   string
	PhotoBarang string
	TimeCreated time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type UserOtlet struct {
	KodeUser     string `gorm:"primary_key"`
	NamaLengkap  string
	NomorKtp     string
	TanggalLahir string
	JenisKelamin string
	FotoKtp      string
	SelfieKtp    string
	FotoKontrak  string
	FotoToko     string
	Email        string
	Hp           string `gorm:"unique;not null"`
	Status       string
	TimeCreated  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Toko struct {
	KodeUser      string
	NamaToko      string
	AlamatToko    string
	KodeReferensi string
	FotoToko      string
	TimeCreated   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type OtpCode struct {
	KodeUser string
	KodeOtp  string
	Hp       string `gorm:"unique;not null"`
}

type Kategory struct {
	KodeKategory int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	IDBarang     string
	KodeUser     string
	NamaKategory string
	TimeCreated  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Pelanggan struct {
	IDPelanggan  string `gorm:"PRIMARY_KEY"`
	KodeUser     string
	Nama         string
	NoTelepon    string
	Email        string
	TanggalLahir time.Time
	JenisKelamin string
	Alamat       string
	Provinsi     string
	KotaKab      string
	Kecamatan    string
	Kelurahan    string
	KodePos      string
	TimeCreated  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Stok struct {
	IdStok      string `gorm:"PRIMARY_KEY"`
	IdKategori  int
	KodeUser    string
	IDBarang    int
	StokAwal    int
	StokAlarm   int
	Pembelian   int
	Penjualan   int
	Penyesuaian int
	StokAkhir   int
	TimeCreated time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Suplier struct {
	IdSuplier   string `gorm:"PRIMARY_KEY"`
	KodeUser    string
	Nama        string
	NoTlp       string
	Email       string
	Alamat      string
	Provinsi    string
	KotaKab     string
	Kecamatan   string
	Kelurahan   string
	KodePos     string
	TimeCreated time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Transaksi struct {
	IdTransaksi    string
	IdBarang       int
	IdSuplier      int
	IdPelanggan    int
	IdPajak        int
	IdDiskon       int
	IdHutang       int
	KodeStruk      string
	JenisTransaksi string
	TipeTransaksi  string
	Qty            int
	Total          int
	Pembulatan     int
	CreateDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
type RincianBelanja struct {
	Rincian           []Rincianbelanja
	TotalPembelanjaan float32
}
type State struct {
	*sync.Mutex                   // inherits locking methods
	Vals        map[string]string // map ids to values
}

type RequestPembelian struct {
	IdBarang string
	Jumlah   int
}

type Rincianbelanja struct {
	NamaBarang  string
	Jumlah      int
	HargaSatuan float32
	HargaTotal  float32
}

type ResponseRegister struct {
	User     UserOtlet
	TokoUser Toko
}

type ResponseOtp struct {
	Messages string
	User     UserOtlet
	TokoUser Toko
}

type ResponseOtpNull struct {
	Messages string
	User     interface{}
	TokoUser interface{}
}

type ResponAddBarang struct {
	Barang     BarangOtlet
	StokBarang Stok
}

type ResponseUniversal struct {
	Message string
	Status  int
}

type Stokbarang struct {
	StokAlarm int
	StokAwal  int
}
