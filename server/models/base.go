package models

type Base struct {
	fields map[string]interface{}
}

func (b *Base) get(key string) interface{} {
	return b.fields[key]
}
