package config

import (
	"flag"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var configurator sync.Once

// GetEnv returns ENV variable from environment or .env file as []byte
func GetEnv(key string) []byte {
	configurator.Do(func() {
		_ = godotenv.Load()
	})
	return []byte(os.Getenv(key))
}

// GetEnvStr returns ENV variable from environment or .env file as string
func GetEnvStr(key string) string {
	return string(GetEnv(key))
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
	mode := gin.ReleaseMode
	if envMode := GetEnvStr("APP_MODE"); envMode != "" {
		mode = envMode
	}
	return flag.String("mode", mode, "Example: -mode=debug")
}
