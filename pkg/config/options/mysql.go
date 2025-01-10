package options

import (
	"time"

	"github.com/spf13/pflag"
	"gorm.io/gorm/logger"
)

type MysqlOptions struct {
	Host     string `json:"host,omitempty" form:"host" mapstructure:"host"`
	Port     int    `json:"port,omitempty" form:"port" mapstructure:"port"`
	Username string `json:"username,omitempty" form:"username" mapstructure:"username"`
	Password string `json:"password,omitempty" form:"password" mapstructure:"password"`
	Database string `json:"database,omitempty" form:"database" mapstructure:"database"`

	MaxIdleConns    int           `json:"maxIdleConns,omitempty" form:"maxIdleConns" mapstructure:"maxIdleConns"`
	MaxOpenConns    int           `json:"maxOpenConns,omitempty" form:"maxOpenConns" mapstructure:"maxOpenConns"`
	ConnMaxLifetime time.Duration `json:"connMaxLifetime,omitempty" form:"connMaxLifetime" mapstructure:"connMaxLifetime"`

	LogLevel int              `json:"logLevel,omitempty" form:"logLevel" mapstructure:"logLevel"`
	Logger   logger.Interface `json:"-" form:"-" mapstructure:"-"`
}

func NewMysqlOptions() *MysqlOptions {
	return &MysqlOptions{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",

		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 300,

		LogLevel: 0,
		Logger:   nil,
	}
}

func (o *MysqlOptions) AddFlags(fs *pflag.FlagSet) {
	pflag.String("mysql.host", "127.0.0.1", "mysql host")
	pflag.Int("mysql.port", 3306, "mysql port")
	pflag.String("mysql.username", "root", "mysql username")
	pflag.String("mysql.password", "root", "mysql password")
	pflag.String("mysql.database", "test", "mysql database")

	pflag.Int("mysql.maxIdleConns", 10, "mysql max idle connections")
	pflag.Int("mysql.maxOpenConns", 100, "mysql max open connections")
	pflag.Int("mysql.connMaxLifetime", 300, "mysql connection max lifetime")

	pflag.Int("mysql.logLevel", 0, "mysql log level")
}
