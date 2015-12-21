package game_data

type MissionGroup struct {
	*Base
}

func (m * MissionGroup) ForClient() map[string]interface{}{
	return map[string]interface{}{
		"id": m.Id,
		"key": m.Key,
	}
}

