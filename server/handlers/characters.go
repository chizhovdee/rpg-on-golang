package handlers
import (
	"github.com/gin-gonic/gin"
	"github.com/chizhovdee/rpg/server/models"
	"fmt"
)

// выполняется при первой загрузке
func CharacterGameData(c *gin.Context){
	character := c.MustGet("current_character").(*models.Character)

	if character == nil {
		// TO DO
		return
	}

	responseEvent(c, "character_game_data_loaded", gin.H{
		"character": character.ForClient(),
	})

	respond(c)
}

// статус персонажа
func CharacterStatus(c *gin.Context) {
	character := c.MustGet("current_character").(*models.Character)

	if character == nil {
		// TO DO
		return
	}

	responseEvent(c, "character_status_loaded", gin.H{
		"character": character.ForClient(),
	})

	respond(c)
}