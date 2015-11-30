package models

type Character struct {
	Base
	Ep int64
	Energy int64
	Hp int64
	Health int64
	Experience int64
}

func (c *Character) AsJson() map[string]interface{} {
	return map[string]interface{}{
		"id": c.Id,
		"ep": c.Ep,
		"energy": c.Energy,
		"hp": c.Hp,
		"health": c.Health,
		"experience": c.Experience,
	}
}
