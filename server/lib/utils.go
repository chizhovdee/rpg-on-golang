package lib
import (
	"log"
	"errors"
)

func ToInt64(v interface{}) int64 {
	var result int64

	switch v.(type) {
		case int8:
			result = int64(v.(int8))
		case int16:
			result = int64(v.(int16))
		case int32:
			result = int64(v.(int32))
		case int64:
			result = int64(v.(int64))
		default:
		  log.Println(errors.New("Не опознанный тип"))
	}

	return result
}