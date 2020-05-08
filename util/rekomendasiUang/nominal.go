package rekomendasiUang

import (
	"math"
)

func pecahan(input float64, nominal float64) (float64, float64) {
	var result float64

	if math.Mod(input, nominal) != 0 {
		result = float64(input) - math.Mod(float64(input), nominal)
		input = math.Mod(float64(input), nominal)
	} else {
		result = input
		if nominal == float64(100) && (input/float64(100)) > 9 {
			result = result - input
		}
		if nominal == float64(1000) && (input/float64(1000)) > 9 {
			result = result - input
		}

		if nominal == float64(10000) && (input/float64(10000)) > 9 {
			result = result - input
		}

		if nominal == float64(100000) && (input/float64(100000)) > 9 {
			result = result - input
		}

	}
	return result, input
}

func unique(intSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Nominal(angka float64) []float64 {

	uang5 := float64(100000)
	uang1 := float64(5000)
	uang2 := float64(10000)
	uang3 := float64(20000)
	uang4 := float64(50000)

	var rekom, rekomResult []float64

	input := float64(angka)
	inputp := float64(angka)
	jutaan, input := pecahan(input, float64(1000000))
	ratusanRibuan, input := pecahan(input, float64(100000))
	puluhanRibuan, input := pecahan(input, float64(10000))
	ribuan, input := pecahan(input, float64(1000))
	ratusan, input := pecahan(input, float64(100))

	rekom = append(rekom, jutaan+ratusanRibuan+uang5)

	if ratusan > 0 {
		rekom = append(rekom, jutaan+ratusanRibuan+puluhanRibuan+ribuan+1000)
	} else {
		rekom = append(rekom, jutaan+ratusanRibuan+puluhanRibuan+ribuan)
	}

	if ribuan > 0 && puluhanRibuan < uang3 {
		rekom = append(rekom, jutaan+ratusanRibuan+puluhanRibuan+uang2)
		rekom = append(rekom, jutaan+ratusanRibuan+puluhanRibuan+uang1)
	}
	if ribuan > 0 && puluhanRibuan <= uang4 {
		rekom = append(rekom, jutaan+ratusanRibuan+puluhanRibuan+uang2)
		rekom = append(rekom, jutaan+ratusanRibuan+uang4)
	}
	if puluhanRibuan > uang2 && puluhanRibuan <= uang3 {
		rekom = append(rekom, jutaan+ratusanRibuan+uang2)
		rekom = append(rekom, jutaan+ratusanRibuan+uang3)
	}

	if puluhanRibuan > uang3 && ribuan >= uang1 && ribuan < uang2 {
		rekom = append(rekom, jutaan+ratusanRibuan+puluhanRibuan+uang2)
	}

	if puluhanRibuan > 0 && ribuan > 0 {
		rekom = append(rekom, jutaan+ratusanRibuan+puluhanRibuan+uang2)
		if ribuan <= 5000 {
			rekom = append(rekom, jutaan+ratusanRibuan+puluhanRibuan+uang1)
		}
	}

	if puluhanRibuan > 0 && puluhanRibuan >= uang2 && ribuan == 0 && ratusan == 0 {
		rekom = append(rekom, jutaan+ratusanRibuan+uang4)

	}

	rekom = unique(rekom)
	for _, data := range rekom {

		if data > inputp {
			rekomResult = append(rekomResult, data)
		}
	}
	if jutaan > 0 && ratusanRibuan == 0 && puluhanRibuan == 0 && ribuan == 0 && ratusan == 0 {
		rekomResult = []float64{jutaan}
	}

	if jutaan == 0 && ratusanRibuan > 0 && puluhanRibuan == 0 && ribuan == 0 && ratusan == 0 {
		rekomResult = []float64{ratusanRibuan}
	}
	if jutaan == 0 && ratusanRibuan == 0 && puluhanRibuan > 0 && ribuan == 0 && ratusan == 0 {
		rekomResult = []float64{puluhanRibuan}
	}
	if jutaan == 0 && ratusanRibuan == 0 && puluhanRibuan == 0 && ribuan > 0 && ratusan == 0 {
		rekomResult = []float64{ribuan}
	}

	if jutaan == 0 && ratusanRibuan == 0 && puluhanRibuan == 0 && ribuan == 0 && ratusan > 0 {
		rekomResult = []float64{ratusan}
	}

	if jutaan > 0 && ratusanRibuan > 0 && puluhanRibuan == 0 && ribuan == 0 && ratusan == 0 {
		rekomResult = []float64{jutaan + ratusanRibuan}
	}

	return rekomResult

}
