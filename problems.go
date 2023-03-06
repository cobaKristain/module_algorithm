package modulesalgorithm

import (
	"fmt"
	"math"
	"strings"
)

func BilanganKeren() string {
	// dipersimpel, count <4 aja cukup
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

func LampuDanTombol(jumlahTombol uint8) {
	var condition bool
	for i := 1; i <= int(jumlahTombol); i++ {
		if int(jumlahTombol)%i == 0 {
			condition = !condition
		}
	}
	if condition == false {
		fmt.Println("Saat ini lampu menjadi mati")
	} else {
		fmt.Println("Saat ini lampu menjadi menyala")
	}
}
func SayHello(nama string) {
	fmt.Printf("Hallo %s, selamat belajar golang!\n", nama)
}

func LuasSegitiga(alas float64, tinggi float64) float64 {
	luas := 0.5 * alas * tinggi
	return luas
}
func KonversiNilai(nilai uint8) {
	var nilaiHuruf string
	if nilai > 100 {
		fmt.Println("Anda memasukkan nilai yang salah")
	} else {
		if nilai >= 80 {
			nilaiHuruf = "A"
		} else if nilai >= 64 {
			nilaiHuruf = "B+"
		} else if nilai >= 50 {
			nilaiHuruf = "B"
		} else if nilai >= 35 {
			nilaiHuruf = "C"
		} else {
			nilaiHuruf = "D"
		}
		fmt.Println("nilai anda adalah:", nilaiHuruf)
	}
}
func LuasPermukaanTabung(tinggiTabung, jariJari float64) {
	lp := 2 * math.Phi * jariJari * (tinggiTabung + jariJari)
	fmt.Printf("Luas permukaan tabung adalah %.2f\n", lp)
}
func FaktorBilangan(cekFaktor int) {
	ke := 1
	for i := 1; i <= cekFaktor; i++ {
		if cekFaktor%i == 0 {
			fmt.Printf("Faktor ke-%d adalah %d\n", ke, i)
			ke++
		}
	}
}

func FaktorBilanganMundur(cekFaktor int) {
	ke := 1
	for i := cekFaktor; i > 0; i-- {
		if cekFaktor%i == 0 {
			fmt.Printf("Faktor ke-%d adalah %d\n", ke, i)
			ke++
		}
	}
}

func PrimeNumber(number int) bool {
	//prima > 1
	//n%i <= 0 truenya tidak lebih dari 2
	prima := true
	count := 0
	if number == 1 {
		prima = false
		return prima
	}
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}if count > 2 {
		prima = false
	}

	return prima
}

func Palindrome(input string) bool {
	//bisa cek https://www.youtube.com/watch?v=DXQuiPKl79Y&t=732s
	var normal []byte
	var balik []byte
	for i := 0; i <= len(input)-1; i++ {
		normal = append(normal, input[i]) //normal[i] = input [i] ini error
	}
	for i := len(input) - 1; i >= 0; i-- {
		balik = append(balik, input[i]) //balik[i] = input [-i] ini error
	}
	for i := 0; i < len(normal); i++ {
		if balik[i] != normal[i] {
			return false
		}
	}
	return true
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func FullPrima(number int) bool {
	if !isPrime(number) {
		return false
	}
	for number > 0 {
		digit := number % 10 // kan sisanya adalah digit terakhir jadi kalau digit 0 ya pasti prima
		if !isPrime(digit) { //masukkanlah sisa baginya.
			return false
		}
		number /= 10
	}
	return true
}
func PlayWithAsterix(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		for e := n; e > i; e-- {
			result += (" ")
		}
		for a := 1; a <= i; a++ {
			result += ("* ")
		}
		result += ("\n")
	}
	return result
}
func DrawXYZ(n int) string {
	var result string
	var nilai int
	// your code here
	// loop dulu sebanyakan n kali
	for i := 0; i < n; i++ { //bawah
		for j := 0; j < n; j++ { //kanan
			nilai = i*n + j + 1
			if nilai%3 == 0 {
				result += "X "
			} else if nilai%2 == 1 {
				result += "Y "
			} else {
				result += "Z "
			}
		}
		result += "\n"
	}
	return result
}
func CetakTabelPerkalian(number int) string {
	var result string
	var nilai int
	// your code here
	// loop dulu sebanyakan n kali
	for i := 1; i <= number; i++ { //bawah
		for j := 1; j <= number; j++ { //kanan
			nilai = i * j
			result += strconv.Itoa(nilai) + "\t"
		}
		result += "\n"
	}
	return result
}
func MeanMedian(arrayInput []float64) (float64, float64) {
	// your code here
	panjangData := len(arrayInput)
	var total float64
	var mean float64
	var median float64

	for i := 0; i < panjangData; i++ {
		total += arrayInput[i]
	}
	mean = total / float64(panjangData)
	if panjangData%2 == 0 {
		median = ((arrayInput[(panjangData/2)-1] + arrayInput[(panjangData/2)]) / 2)
	} else {
		median = arrayInput[panjangData/2]
	}
	return mean, median
}
func UbahHuruf(sentence string) string {
	// your code here
	result := ""
	for _, char := range sentence {
		if char >= 'A' && char <= 'Z' {
			newChar := rune((int(char-'A')+10)%26 + 'A')
			result += string(newChar)
		} else {
			result += string(char)
		}
	}
	return result
}

func Pangkat(base, pow int) int {
	var hasil int = 1
	for i := 1; i <= pow; i++ {
		hasil *= base
	}
	return hasil
}

func BilanganPrimaBaik(n int) string {
	var count int = 0
	for i := 1; i <= n; i++ {
		if count >= 2 {
			return "Bukan Bilangan Prima"
		}
		if n%i == 0 {
			count++
		}
	}
	return "Bilangan Prima"
}

func JoinArrayRemoveDuplicate(x, y []string) []string {
	var arrOutput []string
	var tempMap = make(map[string]bool)

	for _, vals := range x {
		if _, ok := tempMap[vals]; !ok { //tanya sama mas jerry
			tempMap[vals] = true
			arrOutput = append(arrOutput, vals)
		}
	}

	for _, vals := range y {
		if _, ok := tempMap[vals]; !ok {
			tempMap[vals] = true
			arrOutput = append(arrOutput, vals)
		}
	}

	return arrOutput
}


func ArrayUnique(arrayA, arrayB []int) []int {
	var arrayC []int //untuk return (nampung value unique)
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
		if a[i] == b[i] { //awalnya !=
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
		/* 		// Jika karakter adalah spasi, tambahkan spasi ke output
		   	if huruf == ' ' {
		   		output += " "
		   		continue
		   		} */

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
