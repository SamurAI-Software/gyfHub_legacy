package main

import "github.com/jmoiron/sqlx"

// DBdriver is for containing database connection
type DBdriver struct {
	DB *sqlx.DB
}

// API is for containing controllers
type API struct {
	DBDriver *DBdriver
}

//SuccessResponse is struct for passing back success message
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

//SignInInput is struct to contain user sign in info
type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//SignUpInput is struct to contain user sign up info
type SignUpInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
