package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	PORT int
)

func Load() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		panic("can not parse PORT value")
	}

	PORT = parsedPort
}

func IsDevelopment() bool {
	environment := strings.ToLower(os.Getenv("ENVIRONMENT"))

	if environment == "dev" || environment == "devel" || environment == "develop" || environment == "development" {
		return true
	}

	return false
}
