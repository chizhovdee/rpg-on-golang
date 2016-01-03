package models
import (
	"time"
)

//  Атрибуты востанавливающиеся по времени

// public

func (c *Character) IsFull(attribute string) bool {
	var result bool

	switch attribute {
	case "hp":
		result = (c.Restorable("hp") >= c.HealthPoints())
	case "ep":
		result = (c.Restorable("ep") >= c.EnergyPoints())
	}

	return result
}

func (c *Character) Restorable(attribute string) int64 {
	var updated_at time.Time
	var total int64
	var current int64

	switch attribute {
	case "hp":
		updated_at = c.GetHpUpdatedAt()
		total = c.HealthPoints()
		current = c.GetHp()
	case "ep":
		updated_at = c.GetEpUpdatedAt()
		total = c.EnergyPoints()
		current = c.GetEp()
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

// сколько единиц востановилось с последнего обновления ресурса
func (c *Character) restoresSinceLastUpdate(attribute string) int64 {
	var updated_at time.Time

	switch attribute {
	case "hp":
		updated_at = c.GetHpUpdatedAt()
	case "ep":
		updated_at = c.GetEpUpdatedAt()
	}

	return int64((time.Now().Unix() - updated_at.Unix()) / c.restoreSeconds(attribute))
}

// длительность востановления одной единицы ресурса
// возвращает количество секунд
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

func (c *Character) timeToRestore(attribute string) int64 {
	var result int64
	var updated_at time.Time

	switch attribute {
	case "hp":
		updated_at = c.GetHpUpdatedAt()
	case "ep":
		updated_at = c.GetEpUpdatedAt()
	}

	if c.IsFull(attribute) {
		result = 0
	} else {
		restoreSeconds := c.restoreSeconds(attribute)

		result = restoreSeconds - ((time.Now().Unix() - updated_at.Unix()) % restoreSeconds)
	}

	return result
}

func (c *Character) restoreBonus(attribute string) float64 {
	// сумма всех элементов
	return float64(0)
}
