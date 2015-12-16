package game_data
import (
	"errors"
)

type missions struct {
	list []*Mission
	Count int
}

var Missions *missions = &missions{}


func (m *missions) All() []*Mission {
	return m.list
}

func (m *missions) AsJson() []interface{}{
	data := []interface{}{}

	for _, mission := range m.list {
		json := map[string]interface{}{
			"id": mission.Id,
			"key": mission.Key,

		}
		data = append(data, json)
	}

	return data
}

func (m *missions) define(key string, block func(obj Baser)){
	for _, mission := range m.list {
		if mission.Key == key {
			panic(errors.New("Mission with key: " + key + " is exist!"))
		}
	}

	obj := &Mission{Base: &Base{}}
	obj.Key = key

	m.list = append(m.list, define(obj, block).(*Mission))
	m.Count = len(m.list)
}


func defineMissions() {
	defineMissionGroups()

	Missions.define("mission_1", func(obj Baser){
		m := obj.(*Mission)

		m.Mission_group_key = "mission_group_1"
	})

	Missions.define("mission_2", func(obj Baser){
		m := obj.(*Mission)

		m.Mission_group_key = "mission_group_1"
	})
}

