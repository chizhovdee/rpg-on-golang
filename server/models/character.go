package models
import "time"

const (
	FULL_REFILL_DURATION = time.Hour * 12
	HP_RESTORE_DURATION  = time.Minute * 1
	EP_RESTORE_DURATION  = time.Minute * 2
)

type Character struct {
	Base
	Level int64
	Energy int64
	Ep int64
	Ep_updated_at time.Time
	Health int64
	Hp int64
	Hp_updated_at time.Time
	Experience int64
	Basic_money int64
	Vip_money int64
}

// public

func (c *Character) AsJson() map[string]interface{} {
	var fullRefillDuration time.Duration = FULL_REFILL_DURATION
	var hpRestoreDuration time.Duration  = HP_RESTORE_DURATION
	var epRestoreDuration time.Duration  = EP_RESTORE_DURATION

	return map[string]interface{}{
		"id": c.Id,
		"level": c.Level,
		"ep": c.Ep,
		"energy": c.EnergyPoints(),
		"hp": c.Hp,
		"health": c.HealthPoints(),
		"experience": c.Experience,
		"basic_money": c.Basic_money,
		"vip_money": c.Vip_money,
		"full_refill_duration": fullRefillDuration.Seconds(),
		"hp_restore_duration": hpRestoreDuration.Seconds(),
		"ep_restore_duration": epRestoreDuration.Seconds(),
		"hp_restore_bonus": c.restoreBonus("hp"),
		"ep_restore_bonus": c.restoreBonus("ep"),
		"hp_updated_at": c.Hp_updated_at.Unix(),
		"ep_updated_at": c.Ep_updated_at.Unix(),

		// for test
		"restorable_hp": c.RestorableHp(),
		"restorable_ep": c.RestorableEp(),
	}
}

func (c *Character) EnergyPoints() int64 {
	return c.Energy
}

func (c *Character) HealthPoints() int64 {
	return c.Health
}
