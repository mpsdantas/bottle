package env

import (
	"os"
	"strconv"
)

func Get[T any](key string, defaults ...T) T {
	raw := os.Getenv(key)

	if raw == "" && len(defaults) > 0 {
		return defaults[0]
	}

	var zero T
	switch any(zero).(type) {
	case string:
		if raw == "" && len(defaults) > 0 {
			return defaults[0]
		}
		return any(raw).(T)

	case int:
		v, err := strconv.Atoi(raw)
		if err != nil && len(defaults) > 0 {
			return defaults[0]
		}
		return any(v).(T)

	case float64:
		v, err := strconv.ParseFloat(raw, 64)
		if err != nil && len(defaults) > 0 {
			return defaults[0]
		}
		return any(v).(T)

	case bool:
		v, err := strconv.ParseBool(raw)
		if err != nil && len(defaults) > 0 {
			return defaults[0]
		}
		return any(v).(T)
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return zero
}
