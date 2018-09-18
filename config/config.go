package config

import (
	"goqor1.0/config/viperConfig"
	"goqor1.0/config/Interface"
)

type Config = Interface.Config

func Instance() Config {
	return  &viperConfig.Config
}
