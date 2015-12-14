package game_data
import "errors"

type MissionGroup struct {
	*Base
}

type missionGroups struct {
	list []*MissionGroup
	Count int
}

var MissionGroups *missionGroups = &missionGroups{}

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

func defineMissionGroups() {
	MissionGroups.define("mission_group_1", func(obj Baser){})
	MissionGroups.define("mission_group_2", func(obj Baser){})
}