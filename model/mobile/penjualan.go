package mobile

type RequestJualan struct {
	KodeUser		string
	Struk           struk
	Barang          []barang
	IdPelanggan     int
	Diskon          float32
	Subtotal        float32
	PPN             float32
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
