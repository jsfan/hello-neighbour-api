package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type sessionKey string
type masterStore string

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

type KeyPair struct {
	PrivateKey string `yaml:"privateKey"`
	PublicKey  string `yaml:"publicKey"`
}

type Config struct {
	Database    DatabaseConfig `yaml:"database"`
	JwtSignKeys KeyPair        `yaml:"JWTKeys"`
}

const SessionKey sessionKey = "session"
const MasterStore masterStore = "store"

func readFile(fileName string) (fileContents []byte, errVal error) {
	fileHandle, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open config file %s: %w", fileName, err)
	}
	return ioutil.ReadAll(fileHandle)
}

func ReadConfig(fileName string) (cfg *Config, errVal error) {
	rawCfg, err := readFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read config file %s: %w", fileName, err)
	}
	var config *Config
	err = yaml.Unmarshal(rawCfg, &config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config %s: %w", fileName, err)
	}
	return config, nil
}
