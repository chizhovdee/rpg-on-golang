package game_data

import "hash/crc32"


func Define(){
	defineMissionGroups()
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

// возвращает игровые данные в нужном для клиента формате
func forClient(keys []string, list []Baser) map[string]interface{} {
	values := [][]interface{}{}

	for _, obj := range list {
		data := []interface{}{}

		clientData := obj.ForClient()

		for _, key := range keys {
			data = append(data, clientData[key])
		}

		values = append(values, data)
	}

	return map[string]interface{}{"keys": keys, "values": values}
}