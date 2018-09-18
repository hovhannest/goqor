package Interface

import "time"

type Config interface {
	// getters
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
	AllSettings() map[string]interface{}

	Sub(key string) Config

	SetDefault(key string, value interface{})

	Set(key string, value interface{})

	Save(string)
	Read(string, string, string)
}
