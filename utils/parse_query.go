package utils

import (
	"net/http"
	"strconv"
)

func ParseQueryInt(r *http.Request, key string, defaultValue int) int {
	queryValue := r.URL.Query().Get(key)
	if queryValue == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(queryValue)
	if err != nil {
		return defaultValue
	}

	return value
}

func ParseQueryString(r *http.Request, key string) string {
	queryValue := r.URL.Query().Get(key)
	return queryValue
}
