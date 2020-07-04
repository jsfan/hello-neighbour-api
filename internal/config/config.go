package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const SessionKey = "session"
const DatabaseConnection = "storage"

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db"`
}

type KeyPair struct {
	PrivateKey string `yaml:"privateKey"`
	PublicKey  string `yaml:"publicKey"`
}

type Config struct {
	Database    DatabaseConfig `yaml:"database"`
	JwtSignKeys KeyPair        `yaml:"JWTKeys"`
}

func readFile(fileName string) (fileContents []byte, errVal error) {
	fileHandle, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not open config file %s: ", fileName))
	}
	return ioutil.ReadAll(fileHandle)
}

func ReadConfig(fileName string) (cfg *Config, errVal error) {
	rawCfg, err := readFile(fileName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not read config file %s: ", fileName))
	}
	var config *Config
	err = yaml.Unmarshal(rawCfg, &config)
	if err != nil {
		return nil, errors.Wrap(err, "Could not unmarshal config: ")
	}
	return config, nil
}
