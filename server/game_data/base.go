package game_data

type Baser interface {
	setIdByKey()
}

type Base struct {
	Id int64
	Key string
}

func (b *Base) setIdByKey(){
	b.Id = idByKey(b.Key)
}
