package game_data

type QuestGroup struct {
	*Base
}

func (g * QuestGroup) ForClient() map[string]interface{}{
	return map[string]interface{}{
		"id": g.Id,
		"key": g.Key,
	}
}

