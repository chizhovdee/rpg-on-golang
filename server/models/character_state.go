package models
import "log"

type CharacterState struct {
	Quests *QuestsState
}

func FindCharacterState(characterId int64) *CharacterState {
	var err error

	state := &CharacterState{}

	err = App.PgxConn.QueryRow("select quests from character_states where character_id=$1 limit 1", characterId).Scan(
		state.Quests,
	)

	if err != nil {
		log.Println(err.Error())
	}

	return state
}
