package mobile

type RequestJualan struct {
	KodeUser        string
	Struk           struk
	Barang          []barang
	IdPelanggan     int
	DiskonNominal   float32
	DiskonDecimal   float32
	Subtotal        float32
	PPNNominal      float32
	PPNDecimal      float32
	Total           float32
	Pembulatan      float32
	UangTunai       float32
	Kembalian       float32
	Utang           float32
	JenisTransaksi  string
	JenisPembayaran string
}

type struk struct {
	No        string
	NamaKasir string
	Tanggal   string
}

type barang struct {
	IdBarang int
	Jumlah   float32
}

type ListPenjualan struct {
	Nama        string
	IDPelanggan int
	Utang       float32
}

type ResponseListPenjualan struct {
	Penjualan    []ListPenjualan
	TotalPiutang TotalUtang
}

type TotalUtang struct {
	TotalPiutang float32
}
