package game_data

type Mission struct {
	*Base
	Mission_group_key string
}

func (m * Mission) ForClient() map[string]interface{}{
	return map[string]interface{}{
		"id": m.Id,
		"key": m.Key,
		"mission_group_key": m.Mission_group_key,
	}
}
