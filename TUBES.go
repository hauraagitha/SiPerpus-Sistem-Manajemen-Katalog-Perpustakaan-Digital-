package main

import "fmt"

const NMAX = 9999

type Buku struct {
	ID int
	Judul string
	Penulis string
	Kategori string
	Tahun int
	Status string
}

type daftarBuku [NMAX]Buku

func main() {
	var A daftarBuku
	var tombol, Pilih string
	var pilihanpetugas, pilihanpengunjung, idHapus int
	var jalan bool
	var n int
	var peran int


	jalan = true
	n = 0
	tampilan(&tombol)

	for jalan{
		fmt.Println("=======================================================")
		fmt.Println("|              Selamat Datang di SiPerpus             |")
		fmt.Println("=======================================================")
		fmt.Println("|                     Pilih Peran                     |")
		fmt.Println("=======================================================")
		fmt.Println("|  1. Petugas   |   2. Pengunjung   |   3. Keluar     |")
		fmt.Println("=======================================================")
		fmt.Print("Pilihan: ")
		fmt.Scan(&peran)

		if peran == 1 {
			for pilihanpetugas != 4 {
				MenuPetugas(&pilihanpetugas)
				switch pilihanpetugas{
				case 1:
					tambahBuku(&A, &n)
				case 2:
					if n == 0 {
						fmt.Println("Data masih kosong")
					}else{
						ubahData(&A, n)
					}
				case 3:
					if n == 0 {
						fmt.Println("Data masih kosong")
					}else{
						fmt.Println("Masuka ID buku yang akan dihapus: ")
						fmt.Scan(&idHapus)
						hapusBuku(&A, &n, idHapus)
					}
				}
			}
		}else if peran == 2 {
			for pilihanpengunjung != 4 {
				MenuPengunjung(&pilihanpengunjung)
				switch pilihanpengunjung{
				case 1:
					if n == 0{
						fmt.Println("Data masih kosong")
					}else{
						tampilBuku(A, n)
						fmt.Print("Urutkan data? (iya/tidak)?")
						fmt.Scan(&Pilih)
					}
					if Pilih == "iya" {
						urutkanData(&A, n)
						tampilBuku(A, n)
					}
				case 2:
					cariBuku(A, n)
				case 3:
					lihatStatistik(A, n)

				}
			}
		}else{
			jalan = false
		}
	}
	fmt.Println("Program Selesai. Terimakasih suda menggunakan SiPerpus! ")
}

func tampilan(tombol *string) {

	fmt.Println("=======================================================")
	fmt.Println("|        Sistem Manajemen Katalog Perpustakaan        |")
	fmt.Println("|                  Digital (SiPerpus)                 |")
	fmt.Println("=======================================================")

	var input bool
	input = false
	for !input {
		fmt.Print("Tekan Tombol X ")
		fmt.Scan(tombol)

		if *tombol == "x" || *tombol == "X" {
			input = true
		} else {
			fmt.Println("Input salah!")
		}
	}
}
func MenuPetugas(pilihan *int) {

	fmt.Println("=======================================================")
	fmt.Println("|                    MENU UTAMA                       |")
	fmt.Println("=======================================================")
	fmt.Println("|  1. Tambah Data Buku Baru                           |")
	fmt.Println("|  2. Ubah Data Buku                                  |")
	fmt.Println("|  3. Hapus Data Buku                                 |")
	fmt.Println("|  4. Keluar                                          |")
	fmt.Println("=======================================================")
	fmt.Print("Pilihan: ")
	fmt.Scan(pilihan)
}
func MenuPengunjung(pilihan *int) {

	fmt.Println("=======================================================")
	fmt.Println("|                    MENU UTAMA                       |")
	fmt.Println("=======================================================")
	fmt.Println("|  1. Lihat & Urutkan Koleksi Buku                    |")
	fmt.Println("|  2. Cari Buku                                       |")
	fmt.Println("|  3. Lihat Statistik Perpustakaan                    |")
	fmt.Println("|  4. Keluar                                          |")
	fmt.Println("=======================================================")
	fmt.Print("Pilihan: ")
	fmt.Scan(pilihan)

}
func sequentialID(A daftarBuku, n int, idCari int) int {
	var i int
	for i = 0; i < n; i++ {
		if A[i].ID == idCari {
			return i
		}
	}
	return -1
}
func sequentialJudul(A daftarBuku, n int, judulCari string) int {
	var i int
	for i = 0; i < n; i++ {
		if A[i].Judul == judulCari {
			return i
		}
	}
	return -1
}
func binaryID(A daftarBuku, n int, idCari int) int {
	var left, right, mid int
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2

		if A[mid].ID == idCari {
			return mid
		} else if A[mid].ID < idCari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
func binaryJudul(A daftarBuku, n int, judulCari string) int {
	var left, right, mid int
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2

		if A[mid].Judul == judulCari {
			return mid
		} else if A[mid].Judul < judulCari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
func tambahBuku(A *daftarBuku, n *int){
	fmt.Println("========================================================")
	fmt.Println("|   		          TAMBAH BUKU                  |")
	fmt.Println("========================================================")
	fmt.Println("|                      PETUNJUK:                       |")
	fmt.Println("| Gunakan tanda '_' untuk mengganti spasi antar kata   |")
	fmt.Println("========================================================")
	fmt.Println("")
	fmt.Print("ID BUKU       : ")
	fmt.Scan(&A[*n].ID)
	fmt.Print("Judul         : ")
	fmt.Scan(&A[*n].Judul)
	fmt.Print("Penulis       : ")
	fmt.Scan(&A[*n].Penulis)
	fmt.Print("Kategori      : ")
	fmt.Scan(&A[*n].Kategori)
	fmt.Print("Tahun Terbit  : ")
	fmt.Scan(&A[*n].Tahun)
	fmt.Print("Status tersedia (Tersedia/Tidak tersedia): ")
	fmt.Scan(&A[*n].Status)

	*n++
}
func ubahData(A *daftarBuku, n int) {
    var idCari, i int

    fmt.Print("ID buku yang ingin diubah: ")
    fmt.Scan(&idCari)

    i = sequentialID(*A, n, idCari)

    if i != -1 {
        var pilihan int
        var selesai bool
		selesai = false
        for !selesai {
			fmt.Println("=======================================================")
            fmt.Println("|                  MENU UBAH DATA                     |")
			fmt.Println("=======================================================")
            fmt.Println("|  1. Ubah Judul                                      |")
            fmt.Println("|  2. Ubah Penulis                                    |")
            fmt.Println("|  3. Ubah Kategori                                   |")
			fmt.Println("|  4. Ubah Tahun                                      |")
            fmt.Println("|  5. Ubah Status                                     |")
            fmt.Println("|  6. Kembali ke Menu Utama                           |")
			fmt.Println("=======================================================")
            fmt.Print("Pilihan: ")
            fmt.Scan(&pilihan)

            switch pilihan {
            case 1:
                fmt.Print("Masukkan Judul Baru: ")
                fmt.Scan(&A[i].Judul)
				fmt.Scanln()
            case 2:
                fmt.Print("Masukkan Penulis Baru: ")
                fmt.Scan(&A[i].Penulis)
				fmt.Scanln()
			case 3:
                fmt.Print("Masukkan Kategori Baru: ")
                fmt.Scan(&A[i].Kategori)
				fmt.Scanln()
			case 4:
                fmt.Print("Masukkan Tahun Baru: ")
                fmt.Scan(&A[i].Tahun)
				fmt.Scanln()
            case 5:
                fmt.Print("Masukkan Status Baru: ")
                fmt.Scanln(&A[i].Status)
	            fmt.Scanln()
            case 6:
                selesai = true
            }
        }
    } else {
        fmt.Println("Buku tidak ditemukan.")
    }
}
func hapusBuku(A *daftarBuku, n *int, idCari int) {
    var selesai bool
	var i int
	selesai = false

    i  = sequentialID(*A, *n, idCari)

    if i == -1 {
        fmt.Println("Buku tidak ditemukan")
        selesai = true
    }

    if !selesai {
		var j int
        for j = i; j < *n-1; j++ {
            A[j] = A[j+1]
        }

        *n = *n - 1
        fmt.Println("Data dihapus")
        selesai = true
    }
}
func tampilBuku(A daftarBuku, n int) {
	var i int
    if n == 0 {
        fmt.Println("Data kosong")
    } else {
        for i = 0; i < n; i++ {
			fmt.Println(i+1, ".")
			fmt.Println("=======================================================")
            fmt.Println("  ID            : ", A[i].ID)
            fmt.Println("  Judul         : ", A[i].Judul)
            fmt.Println("  Penulis       : ", A[i].Penulis)
            fmt.Println("  Kategori      : ", A[i].Kategori)
			fmt.Println("  Tahun Terbit  : ", A[i].Tahun)
            fmt.Println("  Status        : ", A[i].Status)
			fmt.Println("=======================================================")
        }
    }
}
func selectionSortTahun(A *daftarBuku, n int, arah int) {
	var i, j, idx int
	var temp Buku
    for i = 0; i < n-1; i++ {
        idx = i
        for j = i + 1; j < n; j++ {
            if arah == 1 {
                if A[j].Tahun < A[idx].Tahun {
					idx = j
				}
            } else {
                if A[j].Tahun > A[idx].Tahun {
					idx = j
				}
            }
        }
        temp = A[i]
        A[i] = A[idx]
        A[idx] = temp
    }
}

func insertionSortTahun(A *daftarBuku, n int, arah int) {
    var pass, i int
    var temp Buku

    for pass = 1; pass < n; pass++ {
        temp = A[pass]
        i = pass

        if arah == 1 {
            for i > 0 && A[i-1].Tahun > temp.Tahun {
                A[i] = A[i-1]
                i--
            }
        } else {
            for i > 0 && A[i-1].Tahun < temp.Tahun {
                A[i] = A[i-1]
                i--
            }
        }
        A[i] = temp
    }
}
func urutkanData(A *daftarBuku, n int) {
    var metode, arah int
    fmt.Print("Pilih Metode (1: Selection, 2: Insertion): ")
    fmt.Scan(&metode)
    fmt.Print("Pilih Urutan (1: Ascending, 2: Descending): ")
    fmt.Scan(&arah)

    if metode == 1 {
        selectionSortTahun(A, n, arah)
    } else {
        insertionSortTahun(A, n, arah)
    }
}
func cariBuku(A daftarBuku, n int) {
    var kategori, cara, idCari, i int
	var judulCari string
	fmt.Println("=======================================================")
    fmt.Println("|                   MENU CARI BUKU                    |")
	fmt.Println("=======================================================")
	fmt.Println("|  1. Berdasarkan ID      |   2. Berdasarkan Judul    |")
	fmt.Println("=======================================================")
	fmt.Print("Pilih kategori (1/2): ")
	fmt.Scan(&kategori)
	fmt.Println("=======================================================")
	fmt.Println("|  1. Sequential          |   2. Binary               |")
	fmt.Println("=======================================================")
	fmt.Print("Pilih cara (1/2): ")
    fmt.Scan(&cara)

    if kategori == 1 {
		fmt.Print("Masukkan ID yang dicari: ")
		fmt.Scan(&idCari)
		if cara == 1 {
			i = sequentialID(A, n, idCari)
		}else{
			i = binaryID(A, n, idCari)
		}
    } else {
		fmt.Print("Masukkan judul yang dicari: ")
		fmt.Scan(&judulCari)
		if cara == 1 {
			i = sequentialJudul(A, n, judulCari)
		}else{
			i = binaryJudul(A, n, judulCari)
		}
    }

    if i != -1 {
        fmt.Println("Buku ditemukan!")
		fmt.Println("=======================================================")
        fmt.Println("Judul         : ", A[i].Judul)
        fmt.Println("Penulis       : ", A[i].Penulis)
		fmt.Println("Kategori      : ", A[i].Kategori)
		fmt.Println("Tahun Terbit  : ", A[i].Tahun)
        fmt.Println("Status        : ", A[i].Status)
		fmt.Println("=======================================================")
    } else {
        fmt.Println("Buku tidak ditemukan.")
    }
}
func bukuPerKategori(A daftarBuku, n int) {
	fmt.Println("Statistik per Kategori:")
	var i, j, k int
	var sudahDihitung bool

	for i = 0; i < n; i++ {
		sudahDihitung = false
		for j = 0; j < i; j++ {
			if A[i].Kategori == A[j].Kategori {
				sudahDihitung = true
			}
		}
		if !sudahDihitung {
			var jumlah int
			jumlah = 0
			for k = 0; k < n; k++ {
				if A[k].Kategori == A[i].Kategori {
					jumlah++
				}
			}
			fmt.Printf("- %s  : %d buku\n", A[i].Kategori, jumlah)
		}
	}
	fmt.Println("=======================================================")
}
func lihatStatistik(A daftarBuku, n int) {
    var tersedia, i int
	tersedia = 0

    for i = 0; i < n; i++ {
        if A[i].Status == "Tersedia" {
            tersedia++
        }
    }
	fmt.Println("=======================================================")
    fmt.Println("|               STATISTIK PERPUSTAKAAN                |")
	fmt.Println("=======================================================")
    fmt.Printf("Total Koleksi Buku    : %d\n", n)
    fmt.Printf("Total Koleksi Tersedia: %d\n", tersedia)
	bukuPerKategori(A, n)
}