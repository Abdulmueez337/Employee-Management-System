package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func UserIdExtract(tokenAuth string) (string, int) {
	var MySignKey = []byte("123123123123")
	claims := jwt.MapClaims{}
	tokenString := strings.Split(tokenAuth, " ")[1]
	//fmt.Println("UserId IS: " + tokenString)
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		//return nil,nil
		return MySignKey, nil
	})
	if err != nil {
		fmt.Println("Token Func Error")
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return "", 500
	}
	claimUserId := claims["userID"]
	convertedUserID := fmt.Sprintf("%f", claimUserId)
	tokenUserId := strings.Split(convertedUserID, ".")[0]
	fmt.Println("Token UserId", tokenUserId)
	return tokenUserId, 200
}
