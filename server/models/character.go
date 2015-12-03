package models

type Character struct {
	Base
	Level int64
	Ep int64
	Energy int64
	Hp int64
	Health int64
	Experience int64
	Basic_money int64
	Vip_money int64
}

func (c *Character) AsJson() map[string]interface{} {
	return map[string]interface{}{
		"id": c.Id,
		"level": c.Level,
		"ep": c.Ep,
		"energy": c.Energy,
		"hp": c.Hp,
		"health": c.Health,
		"experience": c.Experience,
		"basic_money": c.Basic_money,
		"vip_money": c.Vip_money,
	}
}
