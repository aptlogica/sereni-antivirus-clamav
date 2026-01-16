package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Antivirus *AntivirusConfig
	// MaxUploadSizeBytes caps total request payload (in bytes)
	MaxUploadSizeBytes int64
}

type AntivirusConfig struct {
	Driver string
	ClamAV ClamAVConfig
}

type ClamAVConfig struct {
	Address        string
	TimeoutSeconds int
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	return &Config{
		Port: GetEnv("PORT", "8080"),
		Antivirus: &AntivirusConfig{
			Driver: GetEnv("ANTIVIRUS_DRIVER", "clamav"),
			ClamAV: ClamAVConfig{
				Address:        GetEnv("CLAMAV_ADDRESS", "127.0.0.1:3310"),
				TimeoutSeconds: GetEnvAsInt("CLAMAV_TIMEOUT_SECONDS", 30),
			},
		},
		MaxUploadSizeBytes: int64(GetEnvAsInt("MAX_UPLOAD_SIZE_MB", 32)) * 1024 * 1024,
	}, nil
}

func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetEnvAsInt(key string, defaultVal int) int {
	valueStr := GetEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
