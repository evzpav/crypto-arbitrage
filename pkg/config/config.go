package config

import (
	"io/ioutil"
	"os"
	"time"

	goex "github.com/nntaoli-project/GoEx"
	builder "github.com/nntaoli-project/GoEx/builder"
	"gopkg.in/yaml.v2"
)

type Config struct {
	APIBuilder      *builder.APIBuilder
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
	apiBuilder := builder.NewAPIBuilder().HttpTimeout(5 * time.Second)
	configs.APIBuilder = apiBuilder
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

//InitExchange returns pointer to exchange client
func (c *Config) InitExchange(name string) goex.API {
	pubkey, seckey := c.GetKeys(name)
	return c.APIBuilder.APIKey(pubkey).APISecretkey(seckey).Build(name)
}

//GetExchangeWrappers get wrappers based on slice of names
func (c *Config) GetExchangeWrappers(exchanges []string) []goex.API {
	wrappers := make([]goex.API, len(exchanges))
	for i, ex := range exchanges {
		wrappers[i] = c.InitExchange(ex)
	}
	return wrappers
}
