package authuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"class-reminder-be/repository"
	"fmt"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func Login(loginData model.LoginRequest) (string, string) {
	// get data tbl_user by username
	data, err := database.GetUserFromDB(loginData.Username)
	if err != nil {
		return "", "01"
	}
	fmt.Println("data", data)
	hashPassword := data["password"].([]byte)
	hashPasswordString := string(hashPassword)
	fmt.Println("===========================================")
	fmt.Println("hashPassword", hashPassword)
	// compare string password with password hash db
	err = bcrypt.CompareHashAndPassword([]byte(hashPasswordString), []byte(loginData.Password))
	if err != nil {
		// Passwords not match
		return "", "02"
	}
	// generate token
	token, err := repository.GenerateToken(loginData.Username)
	fmt.Println("generate token", token)
	if err != nil {
		return "", "01"
	}
	// update to tbl_user
	err = database.UpdateJwtToDB(loginData.Username, token)
	if err != nil {
		fmt.Println(err)
		return "", "01"
	}
	// retrieve token
	return token, ""

}
