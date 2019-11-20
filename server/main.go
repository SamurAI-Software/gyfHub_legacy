package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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

	// fmt.Println(dbPort,
	// 	dBHost,
	// 	dBUser,
	// 	dBPass,
	// 	dBName,
	// 	serverPort,
	// 	intDbPort)

	DBPort := flag.Int("db-port", intDbPort, "DBPort")
	DBHost := flag.String("db-host", dBHost, "DBHost")
	DBUser := flag.String("db-user", dBUser, "DBUser")
	DBPass := flag.String("db-pass", dBPass, "DBPass")
	DBName := flag.String("db-name", dBName, "DBName")
	ServerPort := flag.String("server-port", serverPort, "Port Server Listen")

	flag.Parse()

	// connect to db
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		*DBHost, *DBPort, *DBUser, *DBPass, *DBName)

	conn, err := sqlx.Open("postgres", dbInfo)

	if err != nil {
		panic(err)
	}
	_, err = sqlx.Connect("postgres", dbInfo)

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	err = http.ListenAndServe(*ServerPort, nil)
	if err != nil {
		fmt.Printf("SEVER FAILED TO START, MESSAGE: %s", err.Error())
	}

}
