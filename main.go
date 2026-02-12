package main

import "fmt"

func toF(c float32) float32 {
	return 9 / 5 * (c + 32)
}
func toR(c float32) float32 {
	return 0.8 * c
}
func toK(c float32) float32 {
	return c + 273
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
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
		result := toF(float32(inputSuhu))
		fmt.Printf("%d C = %f Fahrenheit\n", inputSuhu, result)
	} else if targetSuhu == 2 {
		result := toR(float32(inputSuhu))
		fmt.Printf("%d C = %f Reamur\n", inputSuhu, result)
	} else if targetSuhu == 3 {
		result := toK(float32(inputSuhu))
		fmt.Printf("%d C = %f Kelvin\n", inputSuhu, result)
	} else {
		panic("ERROR: Nomor target konversi tidak valid")
	}

}
