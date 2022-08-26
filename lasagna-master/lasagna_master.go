package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTime

}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	n := 0
	s := 0.0
	for _, a := range layers {
		if a == "noodles" {
			n = n + 1
		} else if a == "sauce" {
			s = s + 1
		}

	}
	return (50 * n), (s * 0.2)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	res := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		res[i] = (quantities[i] / 2.0) * float64(portions)
	}
	return res
}
