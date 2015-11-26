package config
import (
	"os"
)

// Проверяются переменные окружения
func CheckEnv() {
	env := []string{
		"ENV",
	}

	for _, value := range env {
		if os.Getenv(value) == "" {
			panic("ENV variable not provided: " + value)
		}
	}
}