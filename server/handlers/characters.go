package handlers
import (
	"github.com/gin-gonic/gin"
	"github.com/chizhovdee/rpg/server/models"
	"log"
	"github.com/chizhovdee/rpg/server/game_data"
)

func CharactersGameData(c *gin.Context){
	character := c.MustGet("current_character").(*models.Character)

	m := game_data.Missions.All()[0]

	log.Println()
	log.Println("------------------------------------------")

	log.Println("Mission", m.Id)
    log.Println("------------------------------------------")

	log.Println(game_data.Missions.AsJson())

	if character == nil {
		// TO DO
		return
	}

	responseEvent(c, "character_game_data_loaded", gin.H{
		"character": character.AsJson(),
		"missions": game_data.Missions.All(),
	})

	respond(c)
}