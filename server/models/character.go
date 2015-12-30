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
import (
	"time"
	"log"
	"github.com/chizhovdee/rpg/server/lib"
)

const (
	FULL_REFILL_DURATION = time.Hour * 12
	HP_RESTORE_DURATION  = time.Second * 30
	EP_RESTORE_DURATION  = time.Second * 45
)

type Character struct {
//	Base
//	Level int64
//	Energy int64
//	Ep int64
//	Ep_updated_at time.Time
//	Health int64
//	Hp int64
//	Hp_updated_at time.Time
//	Experience int64
//	Basic_money int64
//	Vip_money int64

	fields map[string]interface{}
}

// public

func FindCharacter(id int64) *Character {
	ch := &Character{fields: map[string]interface{}{}}

	rows, err := App.PgxConn.Query("select * from characters where id=$1 limit 1", id)

	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		values, _ := rows.Values()

		for index, field := range rows.FieldDescriptions() {
			ch.fields[field.Name] = values[index]
		}
	}

	return ch
}

func (c *Character) AsJson() map[string]interface{} {
	var fullRefillDuration time.Duration = FULL_REFILL_DURATION
	var hpRestoreDuration time.Duration  = HP_RESTORE_DURATION
	var epRestoreDuration time.Duration  = EP_RESTORE_DURATION

	return map[string]interface{}{
		"id": c.fields["id"],
		"level": c.fields["level"],
		"ep": c.fields["ep"],
		"energy": c.fields["energy"],
		"hp": c.fields["hp"],
		"health": c.fields["health"],
		"experience": c.fields["experience"],
		"basic_money": c.fields["basic_money"],
		"vip_money": c.fields["vip_money"],
		"full_refill_duration": fullRefillDuration.Seconds(),
		"hp_restore_duration": hpRestoreDuration.Seconds(),
		"ep_restore_duration": epRestoreDuration.Seconds(),
		"hp_updated_at": c.fields["hp_updated_at"].(time.Time).Unix(),
		"ep_updated_at": c.fields["ep_updated_at"].(time.Time).Unix(),

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

//func (c *Character) SetHp(value int64) {
//	c.Hp = c.Restorable("hp") + value
//	c.Hp_updated_at = time.Now()
//}
//
//func (c *Character) SetEp(value int64) {
//	c.Ep = c.Restorable("ep") + value
//	c.Ep_updated_at = time.Now()
//}

func (c *Character) EnergyPoints() int64 {
	return  lib.ToInt64(c.fields["energy"])
}

func (c *Character) HealthPoints() int64 {
	return  lib.ToInt64(c.fields["health"])
}


