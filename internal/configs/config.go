package configs

import (
	"fmt"
	"strings"
	"github.com/spf13/viper"
)

// resource struct represents a backend service for proxying
type resource struct {
	Name            string
	Endpoint        string
	Destination_URL string
}

// configuration struct represents the full configuration, including the server and resources
type configuration struct {
	Server struct {
		Host        string
		Listen_port string
		CertFile    string 
		KeyFile     string
	}
	Resources []resource
}

var Config *configuration

// NewConfiguration loads the configuration from a YAML file and environment variables
func NewConfiguration() (*configuration, error) {
	viper.AddConfigPath("settings")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config file: %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	return Config, nil
}
