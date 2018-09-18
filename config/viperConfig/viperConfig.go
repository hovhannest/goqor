package viperConfig

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goqor1.0/config/Interface"
	"time"
)

type viperConfig struct {
	v *viper.Viper
}

var Config = viperConfig{ v: viper.New() }

func (_config *viperConfig) Get(key string) interface{} {
	return _config.v.Get(key)
}

func (_config *viperConfig) GetBool(key string) bool {
	return _config.v.GetBool(key)
}

func (_config *viperConfig) GetFloat64(key string) float64 {
	return _config.v.GetFloat64(key)
}

func (_config *viperConfig) GetInt(key string) int {
	return _config.v.GetInt(key)
}

func (_config *viperConfig) GetString(key string) string {
	return _config.v.GetString(key)
}

func (_config *viperConfig) GetStringMap(key string) map[string]interface{} {
	return _config.v.GetStringMap(key)
}

func (_config *viperConfig) GetStringMapString(key string) map[string]string {
	return _config.v.GetStringMapString(key)
}

func (_config *viperConfig) GetStringSlice(key string) []string {
	return _config.v.GetStringSlice(key)
}

func (_config *viperConfig) GetTime(key string) time.Time {
	return _config.v.GetTime(key)
}

func (_config *viperConfig) GetDuration(key string) time.Duration {
	return _config.v.GetDuration(key)
}

func (_config *viperConfig) IsSet(key string) bool {
	return _config.v.IsSet(key)
}

func (_config *viperConfig) AllSettings() map[string]interface{} {
	return _config.v.AllSettings()
}

func (_config *viperConfig) Sub(key string) Interface.Config{
	v := _config.v.Sub(key)
	if(v == nil) {
		return nil
	}
	return &viperConfig{_config.v.Sub(key)}
}

func (_config *viperConfig) SetDefault(key string, value interface{})  {
	_config.v.SetDefault(key, value)
}

func (_config *viperConfig) Set(key string, value interface{}) {
	_config.v.Set(key, value)
}

func (_config *viperConfig) Save(fileName string) {
	err := _config.v.WriteConfigAs(fileName)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func (_config *viperConfig) Read(appName string, fileName string, configFormat string) {
	_config.v.SetEnvPrefix(appName)
	_config.v.AutomaticEnv()

	_config.v.SetConfigName(fileName)
	_config.v.SetConfigType(configFormat)
	_config.v.AddConfigPath("/etc/" + appName)   // path to look for the config file in
	_config.v.AddConfigPath("$HOME/." + appName)
	_config.v.AddConfigPath(".")               // optionally look for config in the working directory

	err := _config.v.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		err = _config.v.WriteConfigAs(fileName + "." + configFormat)
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	_config.v.WatchConfig()
	_config.v.OnConfigChange(func(e fsnotify.Event) {
		// TODO: notify app and all appMicroInstances to reconfigure themselves
		fmt.Println("Config file changed:", e.Name)
	})
}