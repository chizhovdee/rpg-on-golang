package models
import (
	"time"
	"gopkg.in/gorp.v1"
)

type Base struct {
	Id         int64
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (c *Base) PreInsert(s gorp.SqlExecutor) error {
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	return nil
}

func (c *Base) PreUpdate(s gorp.SqlExecutor) error {
	c.UpdatedAt = time.Now()
	return nil
}
