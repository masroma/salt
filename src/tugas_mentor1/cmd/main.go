package main

import (
	"fmt"
	"time"
)

type toko struct {
	nama   string
	alamat string
	no     int
}

type product struct {
	namaProduct string
	harga       int
}

type order struct {
	product
	qty       int
	diskon    int
	orderDate time.Time
}

func main() {
	// data toko
	var i toko
	i.nama = "Toko jaya abadi"
	i.alamat = "jl gatot subroto no 5"
	i.no = 7

	// data product
	var p1 product
	p1.namaProduct = "lemari"
	p1.harga = 1900000

	var p2 = product{"sofa", 2300000}
	var p3 = product{"kursi", 800000}

	// data order
	// orderan := order{
	// 	product:   p3,
	// 	qty:       2,
	// 	diskon:    500000,
	// 	orderDate: time.Now(),
	// }

	// data All Order
	var AllOrder = []order{
		{
			product:   p3,
			qty:       2,
			diskon:    100000,
			orderDate: time.Now(),
		},
		{
			product:   p1,
			qty:       2,
			diskon:    100000,
			orderDate: time.Now(),
		},
		{
			product:   p2,
			qty:       1,
			diskon:    0,
			orderDate: time.Now(),
		},
	}

	fmt.Println("Nama toko : ", i.nama)
	fmt.Println("Alamat : ", i.alamat, "No :", i.no)
	fmt.Println("-------------------------------")
	fmt.Println("DAFTAR PRODUCT")
	fmt.Println("product : ", p1.namaProduct, "harga Rp.", p1.harga)
	fmt.Println("product : ", p2.namaProduct, "harga Rp.", p2.harga)
	fmt.Println("product : ", p3.namaProduct, "harga Rp.", p3.harga)
	fmt.Println("-------------------------------")

	for _, orderan := range AllOrder {
		fmt.Println("ORDERAN TANGGAL ", orderan.orderDate)
		fmt.Println("nama product:", orderan.namaProduct)
		fmt.Println("harga:Rp", orderan.harga)
		fmt.Println("qty:", orderan.qty)
		fmt.Println("subtotal:", orderan.harga*orderan.qty)
		fmt.Println("diskon:", orderan.diskon)
		fmt.Println("total", (orderan.harga*orderan.qty)-orderan.diskon)
		fmt.Println("-------------------------------")
	}

}
