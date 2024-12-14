package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	RpcUrl        string `yaml:"rpc_url"`
	Network       string `yaml:"network"`
	Oracle        string `yaml:"oracle"`
	Comptroller   string `yaml:"comptroller"`
	VaiController string `yaml:"vai_controller"`
	Vai           string `yaml:"vai"`
	WBTC          string `yaml:"wbtc"`
	WETH          string `yaml:"weth"`
	vBNB          string `yaml:"vbnb"`
	WBNB          string `yaml:"wbnb"`
	PrivateKey    string `yaml:"private_key"`
	DB            string `yaml:"db"`
	StartHeight   uint64 `yaml:"start_height"`
	Override      bool   `yaml:"override"`
}

// Setup init config
func New(path string) (*Config, error) {
	// config global config instance
	var config = new(Config)
	//h := log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	//log.Root().SetHandler(h)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
