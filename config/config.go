package config

import "os"

var (
	JWTSecret = getEnv("JWT_SECRET", "blog-system-secret-key-2024")
	JWTExpiry = 24 // hours
	Port      = getEnv("PORT", "8080")
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
