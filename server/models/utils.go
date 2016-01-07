package models
import "log"

func FetchFields(query string, values ...interface{}) map[string]interface{} {
	fields := map[string]interface{}{}

	rows, err := App.PgxConn.Query(query, values...)

	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		values, _ := rows.Values()

		for index, field := range rows.FieldDescriptions() {
			fields[field.Name] = values[index]
		}
	}

	return fields
}
