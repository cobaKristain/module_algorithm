package modulesalgorithm

import "fmt"

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
