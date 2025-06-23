package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL    string
	SupabaseURL    string
	SupabaseKey    string
	JWTSecret      string
	Port           string
	Environment    string
}

func LoadConfig() *Config {
	// Load .env file jika ada
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	config := &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		SupabaseURL: getEnv("SUPABASE_URL", ""),
		SupabaseKey: getEnv("SUPABASE_ANON_KEY", ""),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 