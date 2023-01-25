package config

import (
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// GlobalConfig is a read-only global variable.
// Since we use this global variable, once the values of the config file are read and loaded into memory during the server's initialization, any further changes to the config file will not be reflected in this global variable until the server is restarted.
var GlobalConfig Configuration

type Configuration struct {
	// https://github.com/spf13/viper/issues/385#issuecomment-337264721
	Port                int    `mapstructure:"PORT"`
	UseTLS              bool   `mapstructure:"USE_TLS"`
	MasterVersion       int    `mapstructure:"MASTER_VERSION"`
	PackVersion         int    `mapstructure:"PACK_VERSION"`
	IsMaintenanceMode   bool   `mapstructure:"IS_MAINTENANCE_MODE"`
	MainDomain          string `mapstructure:"MAIN_DOMAIN"`
	StorageDomain       string `mapstructure:"STORAGE_DOMAIN"`
	DataStorageEndpoint string `mapstructure:"DATA_STORAGE_ENDPOINT"`
	TLSCertPath         string `mapstructure:"TLS_CERT_PATH"`
	TLSKeyPath          string `mapstructure:"TLS_KEY_PATH"`
	MasterTableFilename string `mapstructure:"MASTER_TABLE_FILENAME"`
	FileListFilename    string `mapstructure:"FILE_LIST_FILENAME"`
}

func LoadConfig(path string) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configNotFoundError viper.ConfigFileNotFoundError
		if ok := errors.As(err, &configNotFoundError); ok {
			log.Fatal().Err(err).Msg("Error reading config file. Please create a config.yml file in the config folder.")
		} else {
			log.Fatal().Err(err).Msg("Error reading config file. Please ensure that the config.yml file is readable.")
		}
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatal().Err(err).Msg("Error decoding config file. Please check that the config.yml file is valid and properly formatted as a YAML file.")
	}
}
