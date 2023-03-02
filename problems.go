package modulesalgorithm

import (
	"fmt"
	"math"
	"strings"
)

func BilanganKeren() string {
	var n int
	var count int
	fmt.Println("Silahkan Masukkan angka yang ingin anda masukkan")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	if count > 4 || count < 0 {
		return "anda memasukkan bilangan Tidak Keren"
	} else {
		return "anda memasukkan bilangan yang keren"
	}
}

func ArrayUnique(arrayA, arrayB []int) []int {
	var arrayC []int //untuk return
	for _, valA := range arrayA {
		unique := true
		for _, valB := range arrayB {
			if valA == valB {
				unique = false
				break
			}
		}
		if unique {
			arrayC = append(arrayC, valA)
		}
	}
	return arrayC
}

func Compare(a, b string) string {
	var kataBaru []string

	// Ambil panjang string yang lebih pendek
	banyakIterasi := len(a)
	if len(b) < banyakIterasi {
		banyakIterasi = len(b)
	}

	// Iterasi sampai panjang string terpendek
	for i := 0; i < banyakIterasi; i++ {
		if a[i] != b[i] {
			kataBaru = append(kataBaru, string(a[i]))
		}
	}
	return strings.Join(kataBaru, "")
}

func FindMaxSumSubArray(k int, arr []int) int {
	if k > len(arr) {
		return -1 //jika nilai k lebih besar dari ukuran array
	}

	maxSum := 0
	windowSum := 0

	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum = windowSum

	for i := k; i < len(arr); i++ {
		windowSum = windowSum - arr[i-k] + arr[i]
		maxSum = int(math.Max(float64(maxSum), float64(windowSum)))
	}
	return maxSum
}

func RemoveDuplicates(array []int) int {
	indeksUnikSelanjutnya := 1 //indeks untuk memasukkan nilai unik selanjutnya

	for i := 1; i < len(array); i++ {
		if array[indeksUnikSelanjutnya-1] != array[i] {
			//jika nilai i tidak sama dengan nilai sebelumnya, maka i akan dimasukkan ke indekUnikSelanjutnya
			array[indeksUnikSelanjutnya] = array[i]
			indeksUnikSelanjutnya++
		}
	}
	return indeksUnikSelanjutnya
}

// Fungsi caesar menerima offset integer dan sebuah string input
// Fungsi ini menghasilkan string baru, di mana setiap hurufnya bergeser sebanyak offset.
// Asumsikan string input hanya berisi huruf kecil dan spasi.
// Ketika huruf 'z' digeser tiga huruf, lingkaran kembali ke awal alfabet untuk menghasilkan huruf 'c'.
// ASCII huruf kecil dimulai dari 97 untuk 'a' hingga 122 untuk 'z'
func Caesar(offset int, input string) string {
	// Membuat variabel output sebagai string kosong
	output := ""

	// Iterasi melalui setiap karakter dalam string input
	for _, huruf := range input {
		// Jika karakter adalah spasi, tambahkan spasi ke output
		if huruf == ' ' {
			output += " "
			continue
		}

		// Menghitung kode ASCII baru dari karakter yang digeser
		newHuruf := int(huruf) + offset

		// Jika kode ASCII baru lebih besar dari kode ASCII 'z', lingkaran kembali ke awal alfabet
		if newHuruf > 122 {
			newHuruf = 96 + (newHuruf % 122)
		}

		// Tambahkan karakter yang digeser ke output
		output += string(newHuruf)
	}

	// Mengembalikan string output hasil penggeseran
	return output
}
