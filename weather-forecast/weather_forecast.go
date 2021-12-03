//Package weather does stuff.
package weather

//CurrentCondition Weather condition.
var CurrentCondition string

//CurrentLocation place.
var CurrentLocation string

// Forecast current location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
