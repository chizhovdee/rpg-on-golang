package game_data

type Baser interface {
	setIdByKey()
	ForClient() map[string]interface{}
}

type Base struct {
	Id int64
	Key string
}

func (b *Base) setIdByKey(){
	b.Id = idByKey(b.Key)
}
