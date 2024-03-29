package controllers

import (
	"auth/models"
	"context"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpUser(ctx *gin.Context) {
	var user_input models.UserSignUp
	if err := ctx.ShouldBindJSON(&user_input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := map[string]string{"firstname": user_input.Firstname, "lastname": user_input.Lastname, "email": user_input.Email, "phone_number": user_input.Phone_number}
	for k, v := range session {
		err := models.Cache.HSet(context.Background(), "user-session:123", k, v).Err()
		if err != nil {
			panic(err)
		}
	}
	Save_token_send_email(user_input.Email, models.USER_SIGN_UP)
	// userSesion := client.HGetAll(context.Background(), "user-session:123").Val()
	TokenProfile := models.TokenProfile{
		Message:  fmt.Sprintf("Email with token sent to %s", user_input.Email),
		Validity: 3600, // Validity in seconds (e.g., 1 hour)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": TokenProfile})

}
