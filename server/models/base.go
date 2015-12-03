package models
import (
	"time"
	"gopkg.in/gorp.v1"
)

type Base struct {
	Id         int64
	Created_at  time.Time
	Updated_at  time.Time
}

func (c *Base) PreInsert(s gorp.SqlExecutor) error {
	c.Created_at = time.Now()
	c.Updated_at = c.Created_at
	return nil
}

func (c *Base) PreUpdate(s gorp.SqlExecutor) error {
	c.Updated_at = time.Now()
	return nil
}
