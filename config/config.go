package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort string
	LogFile    string
	EthRPCURL  string
}

var AppConfig Config

const (
	// Default configuration values
	DefaultServerPort = "8080"
	DefaultLogFile    = "app.log"
	DefaultEthRPCURL  = "https://ethereum-rpc.publicnode.com"
)

func init() {
	// Load .env file if it exists
	_ = godotenv.Load()

	// Load configuration
	AppConfig = Config{
		ServerPort: getEnv("SERVER_PORT", DefaultServerPort),
		LogFile:    getEnv("LOG_FILE", DefaultLogFile),
		EthRPCURL:  getEnv("ETH_RPC_URL", DefaultEthRPCURL),
	}
}

// getEnv fetches environment variables with a default fallback
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
