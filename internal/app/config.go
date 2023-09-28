package app

import (
	"github.com/core-go/core"
	"github.com/core-go/core/cors"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
)

type Config struct {
	Server     core.ServerConf  `mapstructure:"server"`
	Allow      cors.AllowConfig `mapstructure:"allow"`
	DB         string           `mapstructure:"db"`
	Log        log.Config       `mapstructure:"log"`
	MiddleWare mid.LogConfig    `mapstructure:"middleware"`
}
