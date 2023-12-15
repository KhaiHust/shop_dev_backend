package config

type Config struct {
	HttpConfig server    `mapstructure:"server"`
	DB         database  `mapstructure:"database"`
	LogLevel   logConfig `mapstructure:"log"`
}
type database struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	DatabaseName string `mapstructure:"databaseName"`
}
type server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
type logConfig struct {
	Level string `mapstructure:"level"`
}
