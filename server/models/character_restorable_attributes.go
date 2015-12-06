package models
import "time"

//  Атрибуты востанавливающиеся по времени

func (c *Character) RestorableEp() int64 {
	if c.Ep_updated_at.Unix() < 0 {
		return c.Ep

	} else if c.Ep_updated_at.Before(time.Now().Add(-FULL_REFILL_DURATION)) {
		return c.EnergyPoints()

	} else {
		calculatedValue := c.Ep + c.restoresSinceLastUpdate("ep")

		if calculatedValue >= c.HealthPoints() {
			return c.EnergyPoints()

		} else if calculatedValue < 0 {
			return 0
		}

		return calculatedValue
	}
}

func (c *Character) RestorableHp() int64 {
	if c.Hp_updated_at.Unix() < 0 {
		return c.Hp

	} else if c.Hp_updated_at.Before(time.Now().Add(-FULL_REFILL_DURATION)) {
		return c.HealthPoints()

	} else {
		calculatedValue := c.Hp + c.restoresSinceLastUpdate("hp")

		if calculatedValue >= c.HealthPoints() {
			return c.HealthPoints()

		} else if calculatedValue < 0 {
			return 0
		}

		return calculatedValue
	}
}

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
	// для дальнейшего заполнения
	return float64(0)
}
