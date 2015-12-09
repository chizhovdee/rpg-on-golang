package handlers
import (
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
)

func responseEvent(c *gin.Context, eventType string, data gin.H) {
	c.Set("event_type-" + eventType, data)
}

func respond(c *gin.Context) {
	response := []interface{}{}

	for key, data := range c.Keys {
		parsedKey := strings.Split(key, "-")

		if len(parsedKey) == 2 {
			prefix := parsedKey[0]
			eventType := parsedKey[1]

			if prefix == "event_type" {
				response = append(response, gin.H{
					"event_type": eventType,
					"data": data,
				})
			}
		}
	}

	c.JSON(http.StatusOK, response)
}
