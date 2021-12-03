package partyrobot

import "fmt"
import "strings"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {

	welcomeToMyParty := fmt.Sprintf("Welcome to my party, %s!", name)
	tableNumber := fmt.Sprintf("You have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	nextTo := fmt.Sprintf("You will be sitting next to %s.", neighbor)

	joinedMessage := []string{welcomeToMyParty, tableNumber, nextTo}
	return strings.Join(joinedMessage, "\n")

}
