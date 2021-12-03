package purchase

import "strings"
import "sort"
import "fmt"
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if strings.Contains(kind, "truck"){
		return true
	} else if strings.Contains(kind, "car") {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	cars := []string{option1, option2}
	sort.Strings(cars)
	cars = cars[:len(cars)-1]

	return fmt.Sprintf("%s is clearly the better choice.", strings.Join(cars, ""))
	//return "%s is clearly the better choice.", strings.Join(cars, "")
	
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	
	var value float64
	
	if age <= 3 {
		value = 0.8
	} else if age < 10 {
		value = 0.7
	} else {
		value = 0.5
	}
	return float64(originalPrice * value)
}
