package mobile

import "time"

type DetailTransaksi struct {
	KodeStruk      string
	TipeTransaksi  string
	CreateDate		 time.Time

}
type BarangTransaksi struct {
	NamaBarang	string
	HargaJual	float32
	Harga	float32
	Qty float32
}

type ResponseGetStruk struct {
	Detail DetailTransaksi
	Barang []BarangTransaksi
	Pajak          float32
	Diskon         float32
	Subtotal       float32
	Total          float32
	Pembulatan     float32
	Utang          float32

}
/**
select transaksi_barang.kode_user,b.nama_barang,b.harga_jual,transaksi_barang.qty,b.harga_jual*transaksi_barang.qty as harga from barang_otlet as b
inner join transaksi_barang
on b.id_barang = transaksi_barang.id_barang;


ALTER TABLE transaksi_barang
ADD COLUMN kode_user text;
 */
