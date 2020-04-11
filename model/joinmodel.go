package model

type BarangJoinStok struct {
	KodeUser     string
	IdBarang     string
	NamaBarang   string
	NamaKategory string
	HargaJual    float32
	Jumlah       int
}

type KategoriJoinStok struct {
	KodeKategory int
	NamaKategory string
	KodeUser     string
	Jumlah       int
}

type KategoryJStockBrng struct {
	KodeUser     string
	IdBarang     string
	NamaBarang   string
	NamaKategory string
	Penjualan    int
	StokAkhir    int
}
