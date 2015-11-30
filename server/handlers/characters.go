package handlers
import (
	"github.com/gin-gonic/gin"
	"github.com/chizhovdee/rpg/server/models"
	"fmt"
)

func CharactersGameData(c *gin.Context){
	obj, err := App.DbMap.Get(models.Character{}, 1)

	if err != nil {
		fmt.Println(err.Error())
	}

	character := obj.(*models.Character)

	fmt.Println("Health", character.Health)

	responseEvent(c, "character_game_data_loaded", gin.H{
		"character": character.AsJson(),
	})

	respond(c)
}