package handlers
import (
	"github.com/gin-gonic/gin"
	"github.com/chizhovdee/rpg/server/models"
	"log"
)

func CharactersGameData(c *gin.Context){
	character := c.MustGet("current_character").(*models.Character)

	log.Println("Character", character)

	if character == nil {
		return
	}

	responseEvent(c, "character_game_data_loaded", gin.H{
		"character": character.AsJson(),
	})

	respond(c)
}