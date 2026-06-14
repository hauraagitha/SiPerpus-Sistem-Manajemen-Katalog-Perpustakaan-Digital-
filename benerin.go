package main

import "fmt"

const NMAX = 100

type Peserta struct {
	id     string
	nama   string
	nilai  int
	durasi int
}

type TabPeserta [NMAX]Peserta

func lebihBesar(a, b Peserta) bool {
	if a.nilai != b.nilai {
		return a.nilai > b.nilai
	}
	return a.durasi < b.durasi
}

func insertionSort(data *TabPeserta, n int) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && lebihBesar(key, data[j]) {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func main() {
	var n int
	var data TabPeserta
	var totalNilai float64

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].id, &data[i].nama, &data[i].nilai, &data[i].durasi)
		totalNilai = totalNilai + float64(data[i].nilai)
	}

	insertionSort(&data, n)

	fmt.Println("Data setelah diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", data[i].id, data[i].nama, data[i].nilai, data[i].durasi)
	}

	fmt.Printf("\nPeserta terbaik:\n%s %s %d %d\n", data[0].id, data[0].nama, data[0].nilai, data[0].durasi)

	rataRata := totalNilai / float64(n)
	count := 0
	for i := 0; i < n; i++ {
		if float64(data[i].nilai) > rataRata {
			count++
		}
	}
	fmt.Printf("\nJumlah peserta di atas rata-rata: %d\n", count)
}
