package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/joho/godotenv"
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

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword)
	

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
