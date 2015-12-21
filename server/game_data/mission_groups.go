package game_data
import "errors"

type missionGroups struct {
	list []*MissionGroup
	Count int
}

var MissionGroups *missionGroups = &missionGroups{}

func (m *missionGroups) All() []*MissionGroup {
	return m.list
}

func (m *missionGroups) ForClient() map[string]interface{} {
	return forClient(
		[]string{
			"id",
			"key",
		},
		m.baserList(),
	)
}

func (m *missionGroups) baserList() []Baser {
	list := make([]Baser, m.Count)

	for index, obj := range m.All() {
		list[index] = obj
	}

	return list
}

func (m *missionGroups) define(key string, block func(obj Baser)){
	for _, mission := range m.list {
		if mission.Key == key {
			panic(errors.New("Mission Group with key: " + key + " is exist!"))
		}
	}

	obj := &MissionGroup{Base: &Base{}}
	obj.Key = key

	m.list = append(m.list, define(obj, block).(*MissionGroup))
	m.Count = len(m.list)
}