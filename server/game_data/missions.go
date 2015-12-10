package game_data

func LoadMissions() *Mission {
	return Define(&Mission{}, func(obj Baser){
		m := obj.(*Mission)

		m.Base = &Base{Key: "mission_1"}

	}).(*Mission)
}



