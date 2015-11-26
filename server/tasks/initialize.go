package tasks
import (
	"os"
)

func init(){
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "development")
	}
}
