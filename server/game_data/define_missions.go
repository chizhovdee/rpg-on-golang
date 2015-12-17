package game_data

func defineMissions() {
	Missions.define("mission_1", func(obj Baser){
		m := obj.(*Mission)

		m.Mission_group_key = "mission_group_1"
	})

	Missions.define("mission_2", func(obj Baser){
		m := obj.(*Mission)

		m.Mission_group_key = "mission_group_1"
	})
}