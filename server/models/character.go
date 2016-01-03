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
	fields map[string]interface{}
}

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


// метод возвращает данные для клиента
// этот метод используется и для игровых данных и для статуса
func (c *Character) ForClient() map[string]interface{} {
	// поля берем напрямую без преобразования типов
	return map[string]interface{}{
		"id": c.get("id"),
		"level": c.get("level"),
		"restorable_ep": c.Restorable("ep"),
		"energy_points": c.get("energy"),
		"restorable_hp": c.Restorable("hp"),
		"health_points": c.get("health"),
		"experience": c.get("experience"),
		"basic_money": c.get("basic_money"),
		"vip_money": c.get("vip_money"),
		"hp_restore_in": c.timeToRestore("hp"),
		"ep_restore_in": c.timeToRestore("ep"),
	}
}

func (c *Character) GetHpUpdatedAt() time.Time {
	return c.get("hp_updated_at").(time.Time)
}

func (c *Character) GetEpUpdatedAt() time.Time {
	return c.get("ep_updated_at").(time.Time)
}

func (c *Character) GetHp() int64 {
	return  lib.ToInt64(c.get("hp"))
}

func (c *Character) GetEp() int64 {
	return  lib.ToInt64(c.get("ep"))
}

func (c *Character) EnergyPoints() int64 {
	// сумма всех элементов
	return  lib.ToInt64(c.get("energy"))
}

func (c *Character) HealthPoints() int64 {
	// сумма всех элементов
	return lib.ToInt64(c.get("health"))
}

func (c *Character) get(key string) interface{} {
	return c.fields[key]
}
