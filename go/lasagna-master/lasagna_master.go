package lasagna

func PreparationTime(layers []string, preparationTimePerLayer int) int {
	if preparationTimePerLayer == 0 {
		return 4
	}
	return len(layers) * preparationTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.00
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		default:
			continue
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := make([]float64, 0)
	for _, quantity := range quantities {
		newQuantities = append(newQuantities, (quantity/2)*float64(portions))
	}
	return newQuantities
}
