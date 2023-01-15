package config

type Configuration struct {
	// https://github.com/spf13/viper/issues/385#issuecomment-337264721
	Port int `mapstructure:"port"`
}
