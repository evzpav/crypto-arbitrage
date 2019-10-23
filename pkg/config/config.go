package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ExchangeConfigs []struct {
		ExchangeName string `yaml:"exchange"`   // Represents the exchange name.
		PublicKey    string `yaml:"public_key"` // Represents the public key used to connect to Exchange API.
		SecretKey    string `yaml:"secret_key"` // Represents the secret key used to connect to Exchange API.
	} `yaml:"exchange_configs"`
}

func NewConfig(configFilePath string) (*Config, error) {
	var configs Config
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	contentToMarshal, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(contentToMarshal, &configs)
	if err != nil {
		return nil, err
	}
	return &configs, nil
}

func (c *Config) GetKeys(name string) (pubkey, seckey string) {
	exc := c.ExchangeConfigs
	for i := range exc {
		if exc[i].ExchangeName == name {
			return exc[i].PublicKey, exc[i].SecretKey
		}
	}
	return "", ""
}
