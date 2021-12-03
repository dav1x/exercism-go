package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	
customerString := fmt.Sprintf("Welcome to the Tech Palace, %s", strings.Title(strings.ToUpper(customer)))
	return customerString
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	
	messageStringStars := strings.Repeat("*", numStarsPerLine)
	joinedMessage := []string{messageStringStars, welcomeMsg, messageStringStars}
	return strings.Join(joinedMessage, "\n")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	noStarsMsg := strings.ReplaceAll(oldMsg, "*", "")
	
	return strings.TrimSpace(noStarsMsg)
}
