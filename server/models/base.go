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

func (model *Base) PreInsert(s gorp.SqlExecutor) error {
	model.Created_at = time.Now()
	model.Updated_at = model.Created_at
	return nil
}

func (model *Base) PreUpdate(s gorp.SqlExecutor) error {
	model.Updated_at = time.Now()
	return nil
}
