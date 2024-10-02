package utama

import "fmt"

func utama() {

	var num1, num2, num3, num4, num5 int
	var total int
	fmt.Scanln(&num1, &num2, &num3, &num4, &num5)
	total = num1 + num2 + num3 + num4 + num5
	fmt.Println("Hasil dari penjumlahan", num1, num2, num3, num4, num5, "adalah", total)
}
