package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	DatabaseUrl string
	Port        int
	PrettyLogs  bool

	IsDevelopment bool
	IsProduction  bool

	LogsAppPath    string
	LogsAccessPath string
	LogsMaxSize    int
	LogsMaxBackup  int
	LogsMaxAge     int
	LogsCompress   bool
	LogsLocalTime  bool
)

func getBoolEnv(name string) bool {
	v := strings.ToLower(os.Getenv(name))

	if v == "true" || v == "yes" || v == "1" {
		return true
	}

	return false
}

func getIntEnv(name string) (i int, err error) {
	var iStr string

	if iStr = os.Getenv(name); iStr == "" {
		err = fmt.Errorf("%s is empty", name)
		return
	}

	i, err = strconv.Atoi(iStr)

	return
}

func getStringEnv(name string) (string, error) {
	var s string

	if s = os.Getenv(name); s == "" {
		return "", errors.New("")
	}

	return s, nil
}

func Load() (err error) {
	godotenv.Load()

	if DatabaseUrl, err = getStringEnv("DATABASE_URL"); err != nil {
		return
	}

	if Port, err = getIntEnv("PORT"); err != nil {
		return
	}

	PrettyLogs = getBoolEnv("PRETTY_LOGS")

	if environment := strings.ToLower(os.Getenv("ENVIRONMENT")); environment == "dev" || environment == "devel" || environment == "develop" || environment == "development" {
		IsDevelopment = true
	} else {
		IsProduction = true
	}

	if LogsAppPath, err = getStringEnv("LOGS_APP_PATH"); err != nil {
		return
	}

	if LogsAccessPath, err = getStringEnv("LOGS_ACCESS_PATH"); err != nil {
		return
	}

	if LogsMaxSize, err = getIntEnv("LOGS_MAX_SIZE"); err != nil {
		return
	}

	if LogsMaxBackup, err = getIntEnv("LOGS_MAX_BACKUP"); err != nil {
		return
	}

	if LogsMaxAge, err = getIntEnv("LOGS_MAX_AGE"); err != nil {
		return
	}

	LogsCompress = getBoolEnv("LOGS_COMPRESS")
	LogsLocalTime = getBoolEnv("LOGS_LOCAL_TIME")

	return nil
}
