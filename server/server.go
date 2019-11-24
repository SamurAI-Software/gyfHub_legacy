package main

import (
	"log"
	"net/http"
	"os"

	"flag"
	"fmt"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func auther() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte("secret"), nil)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	// get variables from .env file
	dbPort := os.Getenv("DB-PORT")
	dBHost := os.Getenv("DB-HOST")
	dBUser := os.Getenv("DB-USER")
	dBPass := os.Getenv("DB-PASS")
	dBName := os.Getenv("DB-NAME")
	serverPort := os.Getenv("SERVER-PORT")
	intDbPort, err := strconv.Atoi(dbPort)

	DBPort := flag.Int("db-port", intDbPort, "DBPort")
	DBHost := flag.String("db-host", dBHost, "DBHost")
	DBUser := flag.String("db-user", dBUser, "DBUser")
	DBPass := flag.String("db-pass", dBPass, "DBPass")
	DBName := flag.String("db-name", dBName, "DBName")
	ServerPort := flag.String("server-port", serverPort, "Port Server Listen")

	flag.Parse()

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		*DBHost, *DBPort, *DBUser, *DBPass, *DBName)

	conn, err := sqlx.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	_, err = sqlx.Connect("postgres", dbInfo)
	if err != nil {
		log.Fatalln(err)
	}

	dbDriver := &DBdriver{DB: conn}
	api := &API{DBDriver: dbDriver}

	err = http.ListenAndServe(*ServerPort, api.SetupNewRouter())
	if err != nil {
		fmt.Printf("SEVER FAILED TO START, MESSAGE: %s", err.Error())
	}
}
