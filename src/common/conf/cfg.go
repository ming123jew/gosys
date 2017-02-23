package conf

import (
	cfg "github.com/Unknwon/goconfig"
)

func NewCfg()*cfg.ConfigFile {
	var err error
	c, err := cfg.LoadConfigFile("config.ini")
	if err != nil {
		c, err  = cfg.LoadConfigFile("../config.ini")
	}
	return c
}




