package handlers
import (
	"github.com/gin-gonic/gin"
	"github.com/chizhovdee/rpg/server/models"
)

func CharactersGameData(c *gin.Context){
	character := c.MustGet("current_character").(*models.Character)

	if character == nil {
		// TO DO
		return
	}

	responseEvent(c, "character_game_data_loaded", gin.H{
		"character": character.AsJson(),
	})

	respond(c)
}