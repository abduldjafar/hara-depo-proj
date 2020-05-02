package main

import (
	"fmt"
	"math"
)

func pecahan(input float64, nominal float64) (float64, float64) {
	var result float64
	if math.Mod(input, nominal) != 0 {
		result = float64(input) - math.Mod(float64(input), nominal)
		input = math.Mod(float64(input), nominal)
	} else {
		result = input
	}

	return result, input
}
func main() {

	/**
	uang1 := 5000
	uang2 := 10000
	uang3 := 20000
	uang4 := 50000
	uang5 := 100000

	*/

	input := float64(255000)
	jutaan, input := pecahan(input, float64(1000000))
	ratusanRibuan, input := pecahan(input, float64(100000))
	puluhanRibuan, input := pecahan(input, float64(10000))
	ribuan, input := pecahan(input, float64(1000))
	ratusan, input := pecahan(input, float64(100))

	fmt.Println(jutaan)
	fmt.Println(ratusanRibuan)
	fmt.Println(puluhanRibuan)
	fmt.Println(ribuan)
	fmt.Println(ratusan)
}
