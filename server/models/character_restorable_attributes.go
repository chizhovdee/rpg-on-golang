package models
import (
	"time"
	"github.com/chizhovdee/rpg/server/lib"
)

//  Атрибуты востанавливающиеся по времени

// public
func (c *Character) Restorable(attribute string) int64 {
	var updated_at time.Time
	var total int64
	var current int64

	switch attribute {
	case "hp":
		updated_at = c.fields["hp_updated_at"].(time.Time)
		total = c.HealthPoints()
		current = lib.ToInt64(c.fields["hp"])
	case "ep":
		updated_at = c.fields["ep_updated_at"].(time.Time)
		total = c.EnergyPoints()
		current = lib.ToInt64(c.fields["ep"])
	}

	if updated_at.Before(time.Now().Add(-FULL_REFILL_DURATION)) {
		return total

	} else {
		calculatedValue := current + c.restoresSinceLastUpdate(attribute)

		if calculatedValue >= total {
			return total

		} else if calculatedValue < 0 {
			return 0
		}

		return calculatedValue
	}
}

// private

func (c *Character) restoresSinceLastUpdate(attribute string) int64 {
	var updated_at time.Time

	switch attribute {
	case "hp":
		updated_at = c.fields["hp_updated_at"].(time.Time)
	case "ep":
		updated_at = c.fields["ep_updated_at"].(time.Time)
	}

	return int64((time.Now().Unix() - updated_at.Unix()) / c.restoreSeconds(attribute))
}


func (c *Character) restoreSeconds(attribute string) int64 {
	var duration time.Duration

	switch attribute {
	case "hp":
		duration = HP_RESTORE_DURATION
	case "ep":
		duration = EP_RESTORE_DURATION
	}

	return int64(duration.Seconds() * (100 - c.restoreBonus(attribute)) / 100)
}

func (c *Character) restoreBonus(attribute string) float64 {
	return float64(0)
}
