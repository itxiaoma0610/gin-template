package config

type System struct {
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	AES          string `mapstructure:"aes" json:"aes" yaml:"aes"`
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
}
