package config

type MQTT struct {
	Broker string `mapstructure:"broker" json:"broker" yaml:"broker"`
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
}
