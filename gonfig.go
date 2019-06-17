package gonfig

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var configurator sync.Once

// GetEnv returns ENV variable from environment or .env file as []byte if it's possible and ENV variable exists
// Default 0
func GetEnv(key string) []byte {
	configurator.Do(func() {
		_ = godotenv.Load()
	})
	return []byte(os.Getenv(key))
}

// GetEnvStr returns ENV variable from environment or .env file as string if it's possible and ENV variable exists
// Default ""
func GetEnvStr(key string) string {
	return string(GetEnv(key))
}

// GetEnvArrStr returns ENV variable from environment or .env file as []string if it's possible and ENV variable exists
// Default ""
func GetEnvArrStr(key string) []string {
	str := string(GetEnv(key))
	if str == "" {
		return nil
	}
	return strings.Split(str, ";")
}

// GetEnvStrWithDefault returns ENV variable from environment or .env file as string else returns default value
func GetEnvStrWithDefault(key string, defaultValue string) string {
	if value := string(GetEnv(key)); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvInt returns ENV variable from environment or .env file as int if it's possible and ENV variable exists
// Default 0
func GetEnvInt(key string) int {
	result, err := strconv.Atoi(GetEnvStr(key))
	if err != nil {
		log.Println(err)
	}
	return result
}

// GetEnvIntWithDefault returns ENV variable from environment or .env file as int else returns default value
func GetEnvIntWithDefault(key string, defaultValue int) int {
	if result, err := strconv.Atoi(GetEnvStr(key)); err == nil {
		return result
	}
	return defaultValue
}

// GetListenPort returns a flag to the port to launch the application.
// Looks at the PORT environment variable if the application is running without a flag.
// Default 80
func GetListenPort() *string {
	port := "80"
	if envPort := GetEnvStr("PORT"); envPort != "" {
		port = envPort
	}
	return flag.String("port", port, "Example: -port=8080")
}

//GetApplicationMode returns the flag to the application launch mode.
// Looks at the APP_MODE environment variable if the application is running without a flag -mode
// Default release
func GetApplicationMode() *string {
	mode := "release"
	if envMode := GetEnvStr("APP_MODE"); envMode != "" {
		mode = envMode
	}
	return flag.String("mode", mode, "Example: -mode=debug")
}
