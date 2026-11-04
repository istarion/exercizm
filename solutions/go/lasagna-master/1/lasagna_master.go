package lasagnamaster

func PreparationTime(layers []string, t int) int {
    if t == 0 {
        t = 2
    }
    return len(layers) * t
}

func Quantities(layers []string) (int, float64) {
    NOODLE_LAYER := 50
    SAUCE_LAYER := 0.2
    noodle_weight := 0
    sauce_liters := 0.0
    for _, l := range layers {
        switch l {
        case "noodles": noodle_weight += NOODLE_LAYER
        case "sauce": sauce_liters += SAUCE_LAYER
        }
    }
    return noodle_weight, sauce_liters
}

func AddSecretIngredient(friendList, myList []string) {
    myList[len(myList)-1] = friendList[len(friendList)-1] 
}

func ScaleRecipe(amounts []float64, n int) []float64 {
    result := make([]float64, 0, len(amounts))
    for _, a := range amounts {
        result = append(result, a / 2.0 * float64(n))
    }
    return result
}

