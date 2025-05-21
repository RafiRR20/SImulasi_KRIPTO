package main

import (
	"fmt"
	"strings"
)

var portofolio = map[string]float64{"BTC": 1.5, "ETH": 10}
var hargaKripto = map[string]int{"BTC": 1000000, "ETH": 50000}
var hargaSimulasi = map[string]int{}
var riwayat []string

func main() {
	for k, v := range hargaKripto {
		hargaSimulasi[k] = v
	}

	for {
		tampilkanMenu()
		var pilihan int
		fmt.Print("Pilih menu (1-8): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lihatHarga()
		case 2:
			var koin string
			var jumlah float64
			fmt.Print("Koin yang ingin dibeli (BTC/ETH): ")
			fmt.Scanln(&koin)
			fmt.Print("Jumlah yang ingin dibeli: ")
			fmt.Scanln(&jumlah)
			beliKripto(strings.ToUpper(koin), jumlah)
		case 3:
			var koin string
			var jumlah float64
			fmt.Print("Koin yang ingin dijual (BTC/ETH): ")
			fmt.Scanln(&koin)
			fmt.Print("Jumlah yang ingin dijual: ")
			fmt.Scanln(&jumlah)
			jualKripto(strings.ToUpper(koin), jumlah)
		case 4:
			lihatPortofolio()
		case 5:
			lihatRiwayat()
		case 6:
			simulasiHarga()
		case 7:
			resetSimulasi()
		case 8:
			fmt.Println("Keluar dari program. Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanMenu() {
	fmt.Println("\n=== MENU UTAMA ===")
	fmt.Println("1. Lihat Harga Kripto")
	fmt.Println("2. Beli Kripto")
	fmt.Println("3. Jual Kripto")
	fmt.Println("4. Lihat Portofolio")
	fmt.Println("5. Lihat Riwayat")
	fmt.Println("6. Simulasi Harga")
	fmt.Println("7. Reset Harga")
	fmt.Println("8. Keluar")
}

func lihatHarga() {
	fmt.Println("\n--- Harga Kripto ---")
	for koin, harga := range hargaSimulasi {
		fmt.Printf("%s: Rp %d\n", koin, harga)
	}
}

func beliKripto(koin string, jumlah float64) {
	harga := hargaSimulasi[koin]
	total := jumlah * float64(harga)
	portofolio[koin] += jumlah
	riwayat = append(riwayat, fmt.Sprintf("Beli %.4f %s - Rp %.0f", jumlah, koin, total))
	fmt.Printf("Berhasil membeli %.4f %s\n", jumlah, koin)
}

func jualKripto(koin string, jumlah float64) {
	if jumlah > portofolio[koin] {
		fmt.Println("Jumlah tidak mencukupi.")
		return
	}
	harga := hargaSimulasi[koin]
	total := jumlah * float64(harga)
	portofolio[koin] -= jumlah
	riwayat = append(riwayat, fmt.Sprintf("Jual %.4f %s - Rp %.0f", jumlah, koin, total))
	fmt.Printf("Berhasil menjual %.4f %s\n", jumlah, koin)
}

func lihatPortofolio() {
	fmt.Println("\n--- Portofolio ---")
	for koin, jumlah := range portofolio {
		total := jumlah * float64(hargaSimulasi[koin])
		fmt.Printf("%s: %.4f (Rp %.0f)\n", koin, jumlah, total)
	}
}

func lihatRiwayat() {
	fmt.Println("\n--- Riwayat Transaksi ---")
	for _, t := range riwayat {
		fmt.Println(t)
	}
}

func simulasiHarga() {
	fmt.Println("\n--- Simulasi Harga Naik Rp5000 ---")
	for koin := range hargaSimulasi {
		hargaSimulasi[koin] += 5000
	}
	lihatHarga()
}

func resetSimulasi() {
	for k, v := range hargaKripto {
		hargaSimulasi[k] = v
	}
	fmt.Println("Harga dikembalikan ke nilai awal.")
}
