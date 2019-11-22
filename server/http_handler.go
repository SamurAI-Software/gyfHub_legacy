package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/SamurAI-Software/gyfHub/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type API struct {
	DB *sqlx.DB
}

type SuccessResponse struct {
	Success bool
	Err     error
	Message string
}

func (a *API) signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &models.User{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check username exist
	alreadyExist, err := models.Users(qm.Where("username=?", user.Username)).Exists(context.Background(), a.DB)
	if err != nil || alreadyExist {
		errStr := "Username already exist"
		fmt.Println(err)
		errResponse(errStr, w, err)
		return
	}

	// check email
	alreadyExist, err = models.Users(qm.Where("email=?", user.Email)).Exists(context.Background(), a.DB)
	if err != nil || alreadyExist {
		fmt.Println(err)
		errStr := "Email already exist"
		errResponse(errStr, w, err)
		return
	}

	// check mobile
	alreadyExist, err = models.Users(qm.Where("mobile=?", user.Mobile)).Exists(context.Background(), a.DB)
	if err != nil || alreadyExist {
		fmt.Println(err)
		errStr := "Email already exist"
		errResponse(errStr, w, err)
		return
	}

	verifyToken, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return
	}

	resetToken, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return
	}

	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil || alreadyExist {
		fmt.Println(err)
		fmt.Println(alreadyExist)
		return
	}

	newUser := &models.User{
		ID:             uid.String(),
		Username:       user.Username,
		Email:          user.Email,
		Mobile:         user.Mobile,
		PasswordHash:   hashedPass,
		VerifyToken:    verifyToken.String(),
		ResetPassToken: resetToken.String(),
		Verified:       false,
		Avatar:         []byte(""),
	}

	err = newUser.Insert(context.Background(), a.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("NEW USER: %s", newUser)

}

func (a *API) signInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &models.User{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check user exist by username
	userObj, err := models.Users(qm.Where("username=?", user.Username)).One(context.Background(), a.DB)

	// username doesnt exist
	if err != nil {
		errStr := "Username doesn't exist"
		fmt.Println(errStr)
		errResponse(errStr, w, err)
		return
	}

	// check email pass match
	exist, err := models.Users(qm.Where("username=? AND password_hash=?", user.Username, userObj.PasswordHash)).Exists(context.Background(), a.DB)
	if err != nil || exist == false {
		errStr := "Username or password incorrect"
		fmt.Println(err)
		fmt.Println(exist)

		errResponse(errStr, w, err)
		return
	}

	json.NewEncoder(w).Encode(&SuccessResponse{Success: true, Message: "signed in"})

}

func errResponse(msg string, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SuccessResponse{Success: false, Message: msg})
}
