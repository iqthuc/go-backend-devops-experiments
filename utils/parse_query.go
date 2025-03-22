package utils

import (
	"log"
	"net/http"
	"strconv"
)

type Parsable interface {
	int | string | float64
}

func ParseQueryParam[T Parsable](r *http.Request, key string) T {
	var result T
	queryValue := r.URL.Query().Get(key)
	if queryValue == "" {
		return result
	}
	switch any(result).(type) {
	case int:
		val, err := strconv.Atoi(queryValue)
		if err != nil {
			log.Println("error parse int")
			return result
		}
		result = any(val).(T)

	case float64:
		val, err := strconv.ParseFloat(queryValue, 64)
		if err != nil {
			log.Println("error parse float64")
			return result
		}
		result = any(val).(T)

	case string:
		result = any(queryValue).(T)
	}

	return result
}
