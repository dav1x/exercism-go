package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	var successRate float64
	if speed == 0 {
		successRate = 0.0
	} else if speed <= 4 {
		successRate = 1.0
	} else if speed <= 8 {
		successRate = 0.9
	} else if speed >= 9 {
		successRate = 0.77
	}
	return successRate
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {

	var successRate float64
	successRate = SuccessRate(speed)

	productionRate := float64(speed) * float64(221) * successRate
	return productionRate
}

// CalculatsucceProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	// is speed here in minutes

	ratePerHour := CalculateProductionRatePerHour(speed)
	ratePerMinute := int(ratePerHour) / 60
	return int(ratePerMinute)

}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	ratePerHour := CalculateProductionRatePerHour(speed)
	if ratePerHour > limit {
		return limit
	}
	return ratePerHour
}
