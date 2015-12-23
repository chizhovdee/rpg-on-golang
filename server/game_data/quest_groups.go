package game_data
import "errors"

type questGroups struct {
	list []*QuestGroup
	Count int
}

var QuestGroups *questGroups = &questGroups{}

func (q *questGroups) All() []*QuestGroup {
	return q.list
}

func (q *questGroups) ForClient() map[string]interface{} {
	return forClient(
		[]string{
			"id",
			"key",
		},
		q.baserList(),
	)
}

func (q *questGroups) baserList() []Baser {
	list := make([]Baser, q.Count)

	for index, obj := range q.All() {
		list[index] = obj
	}

	return list
}

func (q *questGroups) define(key string, block func(obj Baser)){
	for _, group := range q.list {
		if group.Key == key {
			panic(errors.New("Quest Group with key: " + key + " is exist!"))
		}
	}

	obj := &QuestGroup{Base: &Base{}}
	obj.Key = key

	q.list = append(q.list, define(obj, block).(*QuestGroup))
	q.Count = len(q.list)
}