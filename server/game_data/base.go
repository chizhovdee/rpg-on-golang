package game_data
import "hash/crc32"

type Baser interface {
	setIdByKey()
}

type Base struct {
	Id int64
	Key string
}


func Define(obj Baser, define func(obj Baser)) Baser{
	define(obj)

	obj.setIdByKey()

	return obj
}

func (b *Base) setIdByKey(){
	b.Id = idByKey(b.Key)
}

func idByKey(key string) int64 {
	return int64(crc32.ChecksumIEEE([]byte(key)))
}