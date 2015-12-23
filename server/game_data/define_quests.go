package game_data

func defineQuests() {
	Quests.define("quest_1", func(obj Baser){
		m := obj.(*Quest)

		m.Quest_group_key = "quest_group_1"
	})

	Quests.define("quest_2", func(obj Baser){
		m := obj.(*Quest)

		m.Quest_group_key = "quest_group_1"
	})
}