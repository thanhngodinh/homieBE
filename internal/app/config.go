package app

import (
	"github.com/core-go/core"
	"github.com/core-go/core/client"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/sql"
)

type Config struct {
	Server     core.ServerConf     `mapstructure:"server"`
	Sql        sql.Config          `mapstructure:"sql"`
	DB         string                  `mapstructure:"db"`
	Client     client.ClientConfig `mapstructure:"client"`
	Log        log.Config          `mapstructure:"log"`
	MiddleWare mid.LogConfig       `mapstructure:"middleware"`
	Status     *core.StatusConfig  `mapstructure:"status"`
	Action     *core.ActionConfig  `mapstructure:"action"`
}
