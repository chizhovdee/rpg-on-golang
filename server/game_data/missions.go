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

func (m *missions) ForClient() map[string]interface{} {
	return forClient(
		[]string{
				"id",
				"key",
				"mission_group_key",
			},
		m.baserList(),
	)
}

func (m *missions) baserList() []Baser {
	list := make([]Baser, m.Count)

	for index, obj := range m.All() {
		list[index] = obj
	}

	return list
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
