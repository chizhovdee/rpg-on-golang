/**
Методы и поля синхронизинованные с клиентом.
Любые изменение значения полей или реализации на сервере,
должны быть применены к клиенту

Константы
-----------------------------------------
FULL_REFILL_DURATION
HP_RESTORE_DURATION
EP_RESTORE_DURATION

Поля
------------------------------------------
Level
Energy
Ep
Ep_updated_at
Health
Hp_updated_at
Experience
Basic_money
Vip_money


Методы
-----------------------------------------
EnergyPoints()
HealthPoints()
Restorable()
restoresSinceLastUpdate()
restoreSeconds()
restoreBonus()

 */

package models
import "time"

const (
	FULL_REFILL_DURATION = time.Hour * 12
	HP_RESTORE_DURATION  = time.Second * 30
	EP_RESTORE_DURATION  = time.Second * 45
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
		"energy": c.Energy,
		"hp": c.Hp,
		"health": c.Health,
		"experience": c.Experience,
		"basic_money": c.Basic_money,
		"vip_money": c.Vip_money,
		"full_refill_duration": fullRefillDuration.Seconds(),
		"hp_restore_duration": hpRestoreDuration.Seconds(),
		"ep_restore_duration": epRestoreDuration.Seconds(),
		"hp_updated_at": c.Hp_updated_at.Unix(),
		"ep_updated_at": c.Ep_updated_at.Unix(),

		// for test
		"restorable_hp": c.Restorable("hp"),
		"restorable_ep": c.Restorable("ep"),
	}
}


func (c *Character) GetHp() int64 {
	return c.Restorable("hp")
}

func (c *Character) GetEp() int64 {
	return c.Restorable("ep")
}

func (c *Character) SetHp(value int64) {
	c.Hp = c.Restorable("hp") + value
	c.Hp_updated_at = time.Now()
}

func (c *Character) SetEp(value int64) {
	c.Ep = c.Restorable("ep") + value
	c.Ep_updated_at = time.Now()
}

func (c *Character) EnergyPoints() int64 {
	return c.Energy
}

func (c *Character) HealthPoints() int64 {
	return c.Health
}


