package main

import (
	"fmt"
	"os"
	phonedb "phone-number-normalizer/db"
	"strings"
	"unicode"
	"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	dbUser := os.Getenv("user")
	dbPassword := os.Getenv("password")
	dbHost := os.Getenv("host")
    dbPort := os.Getenv("port")
	dbName := os.Getenv("dbname")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword)
	must(phonedb.Reset("pgx",psqlInfo, dbName))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbName)
	must(phonedb.Migrate("pgx", psqlInfo))

	db, err := phonedb.Open("pgx", psqlInfo)
	must(err)
	defer db.Close()

	err = db.Seed()
	must(err)
	

}

func normalizeNumber(phone string) string {
	var normalizedNumber strings.Builder

	for _, char := range phone {
		if unicode.IsDigit(char) {
			normalizedNumber.WriteRune(char)
		}
	}
	return normalizedNumber.String()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
