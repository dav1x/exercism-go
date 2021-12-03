package lasagna

const OvenTime = 40

// TODO: define the 'OvenTime' constant

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {

	RemainingTime := OvenTime - actualMinutesInOven
	return RemainingTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {

	PrepTime := numberOfLayers * 2
	return PrepTime
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	
	PrepTime := numberOfLayers * 2
	ETime := actualMinutesInOven + PrepTime
	return ETime
}
