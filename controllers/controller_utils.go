package controllers

import (
	"auth/models"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func Save_token_send_email(email string, token_type models.TokenType) {
	token := generateRandom6DigitNumber()
	setToken(email, token, token_type)
	fmt.Printf("%v", token)

}

func setToken(email string, token int, token_type models.TokenType) error {
	key := fmt.Sprintf("token:%s", email)
	tokenStore := models.TokenStore{
		Email:     email,
		Token:     token,
		TokenType: token_type,
	}

	data, err := json.Marshal(tokenStore)
	if err != nil {
		return err
	}
	return models.Cache.Set(context.Background(), key, data, 600).Err()
}

func generateRandom6DigitNumber() int {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator with time

	// Generate a random number between 100000 and 999999 (inclusive)
	randomNumber := rand.Intn(900000) + 100000

	return randomNumber
}
