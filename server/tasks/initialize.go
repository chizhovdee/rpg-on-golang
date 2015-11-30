package tasks
import (
	"os"
)

// gofer tasks

func init(){
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "development")
	}
}
