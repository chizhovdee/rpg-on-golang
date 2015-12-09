package main
import (
	"github.com/gin-gonic/gin"
	"github.com/chizhovdee/rpg/server/models"
)

func SetCurrentCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		character := models.FindCharacter(1)

		c.Set("current_character", character)

		c.Next()
	}
}