package game_data

import "hash/crc32"


func Define(){
	defineMissions()
}

func define(obj Baser, block func(obj Baser)) Baser{
	block(obj)

	obj.setIdByKey()

	return obj
}

func idByKey(key string) int64 {
	return int64(crc32.ChecksumIEEE([]byte(key)))
}