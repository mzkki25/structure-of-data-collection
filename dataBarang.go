// created by Akmal Muzakki, Namira Salsabilla, and Haura Adzkia on 27/01/2023

package main

import (
	"fmt"
	"os"
)

const kapasitas int = 1000

type dataBarang struct {
	nama  string
	kode  string
	stok  int
	harga float64
}

type tabDataBarang struct {
	barang [kapasitas]dataBarang
	total  int
}

func inputData(data *tabDataBarang) {
	data.total = 0
	for true {
		fmt.Print("Masukkan kode barang atau STOP untuk selesai: ")
		fmt.Scan(&data.barang[data.total].kode)
		if data.barang[data.total].kode == "STOP" {
			break
		}
		fmt.Print("Masukkan nama barang: ")
		fmt.Scan(&data.barang[data.total].nama)
		fmt.Print("Masukkan stok barang: ")
		fmt.Scan(&data.barang[data.total].stok)
		fmt.Print("Masukkan harga barang per item: ")
		fmt.Scan(&data.barang[data.total].harga)
		fmt.Println("")
		data.total++
	}
}

func updateHargaData(data *tabDataBarang) {
	for i := 0; i < data.total; i++ {
		data.barang[i].harga *= float64(data.barang[i].stok)
	}
}

func tampilkanData(data *tabDataBarang) {
	fmt.Println("======================================================")
	for i := 0; i < data.total; i++ {
		fmt.Println("Data barang ke:", i+1)
		fmt.Println("kode barang:", data.barang[i].kode)
		fmt.Println("Nama barang:", data.barang[i].nama)
		fmt.Println("stok barang:", data.barang[i].stok)
		fmt.Println("harga barang total:", data.barang[i].harga)
		fmt.Println("")
	}
	fmt.Println("======================================================")
}

func totalStokBarangTerdata(data tabDataBarang) float64 {
	var hasil float64
	for i := 0; i < data.total; i++ {
		hasil += float64(data.barang[i].stok)
	}
	return hasil
}

func totalHargaBarangTerdata(data tabDataBarang) float64 {
	var hasil float64
	for i := 0; i < data.total; i++ {
		hasil += float64(data.barang[i].harga)
	}
	return hasil
}

func hargaPalingTinggi(data tabDataBarang) int {
	Max := 0
	for i := 0; i < data.total; i++ {
		if data.barang[i].harga > data.barang[Max].harga {
			Max = i
		}
	}
	return Max
}

func stokPalingBanyak(data tabDataBarang) int {
	barangMax := 0
	for i := 0; i < data.total; i++ {
		if data.barang[i].stok > data.barang[barangMax].stok {
			barangMax = i
		}
	}
	return barangMax
}

func hargaPalingRendah(data tabDataBarang) int {
	Min := 0
	for i := 0; i < data.total; i++ {
		if data.barang[i].harga < data.barang[Min].harga {
			Min = i
		}
	}
	return Min
}

func stokPalingSedikit(data tabDataBarang) int {
	barangMin := 0
	for i := 0; i < data.total; i++ {
		if data.barang[i].stok < data.barang[barangMin].stok {
			barangMin = i
		}
	}
	return barangMin
}

func cariData(data tabDataBarang, kode string) int {
	for i := 0; i < data.total; i++ {
		if data.barang[i].kode == kode {
			return i
		}
	}
	return -1
}

func tampilkanDataDicari(data tabDataBarang) {
	var kode string

	fmt.Println("Pencarian Kode")
	for true {
		fmt.Print("Kode barang: ")
		fmt.Scan(&kode)
		if cariData(data, kode) != -1 {
			fmt.Println("Nama barang:", data.barang[cariData(data, kode)].nama)
			fmt.Println("Stok barang:", data.barang[cariData(data, kode)].stok)
			fmt.Println("Harga barang:", data.barang[cariData(data, kode)].harga)
			fmt.Println("")
		} else {
			fmt.Println("Data tidak ditemukan")
			break
		}
	}
}

func sortHargaDataBarang(data *tabDataBarang) {
	for i := 0; i < (*data).total-1; i++ {
		max := i
		for j := i + 1; j < (*data).total; j++ {
			if (*data).barang[max].harga > (*data).barang[j].harga {
				max = j
			}
		}
		temp := (*data).barang[i]
		(*data).barang[i] = (*data).barang[max]
		(*data).barang[max] = temp
	}
	fmt.Println("urutan data berdasarkan harga secara menurun")
	for i := 0; i < data.total; i++ {
		fmt.Printf("urutan %d. %s -> %.f\n", i+1, data.barang[i].nama, data.barang[i].harga)
	}
}

func hapusData(data *tabDataBarang) {
	var kode string
	fmt.Println("Hapus data berdasarkan kode")
	for true {
		fmt.Print("Kode barang: ")
		fmt.Scan(&kode)
		if cariData(*data, kode) != -1 {
			(*data).barang[cariData(*data, kode)] = (*data).barang[(*data).total-1]
			(*data).total--
			fmt.Println("Data berhasil dihapus")
			break
		} else {
			fmt.Println("Data tidak ditemukan")
			break
		}
	}
}

func main() {
	var data tabDataBarang
	var pilihan int

	fmt.Println("Input data barang")
	inputData(&data)
	updateHargaData(&data)

	for true {
		fmt.Println("1. Hapus data barang")
		fmt.Println("2. Tampilkan data barang")
		fmt.Println("3. Total stok barang terdata")
		fmt.Println("4. Total harga barang terdata")
		fmt.Println("5. Harga barang terdata paling tinggi")
		fmt.Println("6. Stok barang terdata paling banyak")
		fmt.Println("7. Harga barang terdata paling rendah")
		fmt.Println("8. Stok barang terdata paling sedikit")
		fmt.Println("9. Cari data barang")
		fmt.Println("10. Urutkan data barang berdasarkan harga")
		fmt.Println("11. Keluar")

		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println("")

		switch pilihan {
		case 1:
			hapusData(&data)
		case 2:
			tampilkanData(&data)
		case 3:
			fmt.Println("Total stok barang terdata:", totalStokBarangTerdata(data))
		case 4:
			fmt.Println("Total harga barang terdata:", totalHargaBarangTerdata(data))
		case 5:
			fmt.Println("Harga barang terdata paling tinggi:", data.barang[hargaPalingTinggi(data)].nama)
		case 6:
			fmt.Println("Stok barang terdata paling banyak:", data.barang[stokPalingBanyak(data)].nama)
		case 7:
			fmt.Println("Harga barang terdata paling rendah:", data.barang[hargaPalingRendah(data)].nama)
		case 8:
			fmt.Println("Stok barang terdata paling sedikit:", data.barang[stokPalingSedikit(data)].nama)
		case 9:
			tampilkanDataDicari(data)
		case 10:
			sortHargaDataBarang(&data)
		case 11:
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
		fmt.Println("")
	}
}
