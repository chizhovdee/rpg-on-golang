package game_data

type Quest struct {
	*Base
	Quest_group_key string
}

func (q * Quest) ForClient() map[string]interface{}{
	return map[string]interface{}{
		"id": q.Id,
		"key": q.Key,
		"quest_group_key": q.Quest_group_key,
	}
}
