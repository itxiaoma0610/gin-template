package config

type Redis struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Pass string `mapstructure:"pass" json:"pass" yaml:"pass"`
	Type string `mapstructure:"type" json:"type" yaml:"type"`
	DB   string `mapstructure:"db" json:"db" yaml:"db"`
}
