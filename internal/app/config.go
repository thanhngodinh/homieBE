package app

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Server ServerConf  `mapstructure:"server"`
	Allow  AllowConfig `mapstructure:"allow"`
	DB     string      `mapstructure:"db"`
	Log    LogConfig   `mapstructure:"log"`
}

type ServerConf struct {
	AppId             string         `yaml:"app_id" mapstructure:"app_id" json:"appId,omitempty" gorm:"column:appid" bson:"appId,omitempty" dynamodbav:"appId,omitempty" firestore:"appId,omitempty"`
	Name              string         `yaml:"name" mapstructure:"name" json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Version           string         `yaml:"version" mapstructure:"version" json:"version,omitempty" gorm:"column:version" bson:"version,omitempty" dynamodbav:"version,omitempty" firestore:"version,omitempty"`
	Port              *int64         `yaml:"port" mapstructure:"port" json:"port,omitempty" gorm:"column:port" bson:"port,omitempty" dynamodbav:"port,omitempty" firestore:"port,omitempty"`
	Secure            bool           `yaml:"secure" mapstructure:"secure" json:"secure,omitempty" gorm:"column:secure" bson:"secure,omitempty" dynamodbav:"secure,omitempty" firestore:"secure,omitempty"`
	Log               *bool          `yaml:"log" mapstructure:"log" json:"log,omitempty" gorm:"column:log" bson:"log,omitempty" dynamodbav:"log,omitempty" firestore:"log,omitempty"`
	Monitor           *bool          `yaml:"monitor" mapstructure:"monitor" json:"monitor,omitempty" gorm:"column:monitor" bson:"monitor,omitempty" dynamodbav:"monitor,omitempty" firestore:"monitor,omitempty"`
	CORS              *bool          `yaml:"cors" mapstructure:"cors" json:"cors,omitempty" gorm:"column:cors" bson:"cors,omitempty" dynamodbav:"cors,omitempty" firestore:"cors,omitempty"`
	WriteTimeout      *time.Duration `yaml:"write_timeout" mapstructure:"write_timeout" json:"writeTimeout,omitempty" gorm:"column:writetimeout" bson:"writeTimeout,omitempty" dynamodbav:"writeTimeout,omitempty" firestore:"writeTimeout,omitempty"`
	ReadTimeout       *time.Duration `yaml:"read_timeout" mapstructure:"read_timeout" json:"readTimeout,omitempty" gorm:"column:readtimeout" bson:"readTimeout,omitempty" dynamodbav:"readTimeout,omitempty" firestore:"readTimeout,omitempty"`
	ReadHeaderTimeout *time.Duration `yaml:"read_header_timeout" mapstructure:"read_header_timeout" json:"readHeaderTimeout,omitempty" gorm:"column:readheadertimeout" bson:"readHeaderTimeout,omitempty" dynamodbav:"readHeaderTimeout,omitempty" firestore:"readHeaderTimeout,omitempty"`
	IdleTimeout       *time.Duration `yaml:"idle_timeout" mapstructure:"idle_timeout" json:"idleTimeout,omitempty" gorm:"column:idletimeout" bson:"idleTimeout,omitempty" dynamodbav:"idleTimeout,omitempty" firestore:"idleTimeout,omitempty"`
	MaxHeaderBytes    *int           `yaml:"max_header_bytes" mapstructure:"max_header_bytes" json:"maxHeaderBytes,omitempty" gorm:"column:maxheaderbytes" bson:"maxHeaderBytes,omitempty" dynamodbav:"maxHeaderBytes,omitempty" firestore:"maxHeaderBytes,omitempty"`
	Cert              string         `yaml:"cert" mapstructure:"cert" json:"cert,omitempty" gorm:"column:cert" bson:"cert,omitempty" dynamodbav:"cert,omitempty" firestore:"cert,omitempty"`
	Key               string         `yaml:"key" mapstructure:"key" json:"key,omitempty" gorm:"column:key" bson:"key,omitempty" dynamodbav:"key,omitempty" firestore:"key,omitempty"`
}

type AllowConfig struct {
	Origins            string `yaml:"origins" mapstructure:"origins" json:"origins,omitempty" gorm:"column:origins" bson:"origins,omitempty" dynamodbav:"origins,omitempty" firestore:"origins,omitempty"`
	Methods            string `yaml:"methods" mapstructure:"methods" json:"methods,omitempty" gorm:"column:methods" bson:"methods,omitempty" dynamodbav:"methods,omitempty" firestore:"methods,omitempty"`
	Headers            string `yaml:"headers" mapstructure:"headers" json:"headers,omitempty" gorm:"column:headers" bson:"headers,omitempty" dynamodbav:"headers,omitempty" firestore:"headers,omitempty"`
	Credentials        bool   `yaml:"credentials" mapstructure:"credentials" json:"credentials,omitempty" gorm:"column:credentials" bson:"credentials,omitempty" dynamodbav:"credentials,omitempty" firestore:"credentials,omitempty"`
	MaxAge             *int   `yaml:"max_age" mapstructure:"max_age" json:"maxAge,omitempty" gorm:"column:maxage" bson:"maxAge,omitempty" dynamodbav:"maxAge,omitempty" firestore:"maxAge,omitempty"`
	ExposedHeaders     string `yaml:"exposed_headers" mapstructure:"exposed_headers" json:"exposedHeaders,omitempty" gorm:"column:exposedheaders" bson:"exposedHeaders,omitempty" dynamodbav:"exposedHeaders,omitempty" firestore:"exposedHeaders,omitempty"`
	OptionsPassthrough *bool  `yaml:"options_passthrough" mapstructure:"options_passthrough" json:"optionsPassthrough,omitempty" gorm:"column:optionsPassthrough" bson:"optionsPassthrough,omitempty" dynamodbav:"optionsPassthrough,omitempty" firestore:"optionsPassthrough,omitempty"`
}

type LogConfig struct {
	Level           string           `yaml:"level" mapstructure:"level" json:"level,omitempty" gorm:"column:level" bson:"level,omitempty" dynamodbav:"level,omitempty" firestore:"level,omitempty"`
	Output          string           `yaml:"output" mapstructure:"output" json:"output,omitempty" gorm:"column:output" bson:"output,omitempty" dynamodbav:"output,omitempty" firestore:"output,omitempty"`
	Duration        string           `yaml:"duration" mapstructure:"duration" json:"duration,omitempty" gorm:"column:duration" bson:"duration,omitempty" dynamodbav:"duration,omitempty" firestore:"duration,omitempty"`
	Fields          string           `yaml:"fields" mapstructure:"fields" json:"fields,omitempty" gorm:"column:fields" bson:"fields,omitempty" dynamodbav:"fields,omitempty" firestore:"fields,omitempty"`
	FieldMap        string           `yaml:"field_map" mapstructure:"field_map" json:"fieldMap,omitempty" gorm:"column:fieldmap" bson:"fieldMap,omitempty" dynamodbav:"fieldMap,omitempty" firestore:"fieldMap,omitempty"`
	Map             *logrus.FieldMap `yaml:"map" mapstructure:"map" json:"map,omitempty" gorm:"column:map" bson:"map,omitempty" dynamodbav:"map,omitempty" firestore:"map,omitempty"`
	TimestampFormat string           `yaml:"timestamp_format" mapstructure:"timestamp_format" json:"timestampFormat,omitempty" gorm:"column:timestampformat" bson:"timestampFormat,omitempty" dynamodbav:"timestampFormat,omitempty" firestore:"timestampFormat,omitempty"`
}
