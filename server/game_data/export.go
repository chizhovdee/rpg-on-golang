// здесь прописываем все данные, которые надо экспортировать на клиент

package game_data

func ExportToClient() map[string]interface{} {
	Define()

	// ключи добавлять в единственном числе
	// перед добавление нового ключа, нужно добавить класс в client/js/game_data
	return map[string]interface{} {
		"quest_group": QuestGroups.ForClient(),
		"quest": Quests.ForClient(),

		// здесь может быть новый ключ
	}
}