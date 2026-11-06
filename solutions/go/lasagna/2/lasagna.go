package lasagna

const OvenTime = 40
const LayerPreparationTime = 2

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * LayerPreparationTime
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
