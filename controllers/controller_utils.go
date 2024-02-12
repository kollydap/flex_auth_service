package controllers

import (
	"auth/models"
	"fmt"
	"math/rand"
	"time"
)

func Save_token_send_email(email string, token_type models.TokenType) {
	token := generateRandom6DigitNumber()
	fmt.Printf("%v", token)

}

func generateRandom6DigitNumber() int {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator with time

	// Generate a random number between 100000 and 999999 (inclusive)
	randomNumber := rand.Intn(900000) + 100000

	return randomNumber
}
