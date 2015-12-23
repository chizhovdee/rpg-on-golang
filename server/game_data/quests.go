package game_data
import (
	"errors"
)

type quests struct {
	list []*Quest
	Count int
}

var Quests *quests = &quests{}

func (q *quests) All() []*Quest {
	return q.list
}

func (q *quests) ForClient() map[string]interface{} {
	return forClient(
		[]string{
				"id",
				"key",
				"quest_group_key",
			},
		q.baserList(),
	)
}

func (q *quests) baserList() []Baser {
	list := make([]Baser, q.Count)

	for index, obj := range q.All() {
		list[index] = obj
	}

	return list
}

func (m *quests) define(key string, block func(obj Baser)){
	for _, quest := range m.list {
		if quest.Key == key {
			panic(errors.New("Quest with key: " + key + " is exist!"))
		}
	}

	obj := &Quest{Base: &Base{}}
	obj.Key = key

	m.list = append(m.list, define(obj, block).(*Quest))
	m.Count = len(m.list)
}
