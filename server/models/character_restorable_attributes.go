package models
import "time"

//  Атрибуты востанавливающиеся по времени

// public
func (c *Character) Restorable(attribute string) int64 {
	var updated_at time.Time
	var total int64
	var current int64

	switch attribute {
	case "hp":
		updated_at = c.Hp_updated_at
		total = c.HealthPoints()
		current = c.Hp
	case "ep":
		updated_at = c.Ep_updated_at
		total = c.EnergyPoints()
		current = c.Ep
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
		updated_at = c.Hp_updated_at
	case "ep":
		updated_at = c.Ep_updated_at
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
