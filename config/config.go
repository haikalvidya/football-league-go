package config

import "github.com/spf13/viper"

const (
	DEV        = "dev"
	PRODUCTION = "production"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type ServerConfig struct {
	Address           string `mapstructure:"address"`
	Env               string `mapstructure:"env"`
	BaseURL           string `mapstructure:"base_url"`
	InternalAccessKey string `mapstructure:"internal_access_key"`
}

type DatabaseConfig struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	Name               string `mapstructure:"name"`
	Params             string `mapstructure:"params"`
	MigrationTableName string `mapstructure:"migration_table_name"`
}

type JWTConfig struct {
	Secret                 string `mapstructure:"secret"`
	AccessTokenExpiredHour int    `mapstructure:"access_token_expire_hour"`
	RefreshTokenExpireHour int    `mapstructure:"refresh_token_expire_hour"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func Load(cfgName string, paths ...string) (c *Config, err error) {
	viper.SetConfigName(cfgName)
	viper.SetConfigType("yaml")

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}
