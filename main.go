package main

import (
	"fmt"
	"koda-b6-golang4/convertion"
)

func cobalagi() {
	fmt.Println("\n Ingin mencoba lagi?")
	fmt.Println(" 1. Ya\n 2. Tidak")
	fmt.Println("Masukan pilihan: ")
	var isLanjut int
	fmt.Scan(&isLanjut)
	if isLanjut == 1 {
		main()
	} else {
		fmt.Println("Mengakhiri program...")
	}
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			var opsi int
			fmt.Println(r)
			fmt.Println("1. Kembali ke Input")
			fmt.Println("2. Akhiri Program")
			fmt.Println("Masukan Pilihan: ")
			fmt.Scan(&opsi)
			if opsi == 1 {
				main()
			} else {
				fmt.Println("Mengakhiri program...")
			}
		}
	}()

	var inputSuhu int
	var targetSuhu int

	fmt.Println("== KONVERSI SUHU CELCIUS ==")
	fmt.Print("Masukan nilai suhu: ")
	fmt.Scan(&inputSuhu)
	fmt.Println("Konversikan ke :")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Reamur")
	fmt.Println("3. Kelvin")
	fmt.Print("Masukan nomor target konversi: ")
	fmt.Scan(&targetSuhu)

	if targetSuhu == 1 {
		result := convertion.ToF(float32(inputSuhu))
		fmt.Println("------------------------")
		fmt.Printf("%d C = %.2f Fahrenheit\n", inputSuhu, result)
		fmt.Println("------------------------")
		cobalagi()
	} else if targetSuhu == 2 {
		result := convertion.ToR(float32(inputSuhu))
		fmt.Println("------------------------")
		fmt.Printf("%d C = %.2f Reamur\n", inputSuhu, result)
		fmt.Println("------------------------")
		cobalagi()
	} else if targetSuhu == 3 {
		result := convertion.ToK(float32(inputSuhu))
		fmt.Println("------------------------")
		fmt.Printf("%d C = %.2f Kelvin\n", inputSuhu, result)
		fmt.Println("------------------------")
		cobalagi()
	} else {
		panic("ERROR: Nomor target konversi tidak valid")
	}

}
