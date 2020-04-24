package mobile

import "time"

type Pelanggan struct {
	IDPelanggan  int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	KodeUser     string
	Nama         string
	NoTlp        string
	Email        string
	Alamat       string
	Gender       string
	TanggalLahir string
	TimeCreated  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ResponseDetailPelanggan struct {
	Nama              string
	Hp                string
	PelangganSejak    string
	KunjunganTerakhir string
	Piutang           float32
	JumlahTransaksi   float32
	NilaiTransaksi    float32
	RataTransaksi     float32
}

type PlgganSejak struct {
	PelangganSejak string
}

type KungjnTrkhr struct {
	KunjunganTerakhir string
}
