package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"os"
)

var CFG *Config

type ConfigJson struct {
	Database struct {
		Master string `json:"master"`
		Slave  string `json:"slave"`
	} `json:"database"`
	Server struct {
		Port     string `json:"port"`
		GrpcPort string `json:"grpcPort"`
	} `json:"server"`
}

type Config struct {
	Server         types.ServerConfig
	Database       types.DatabaseConfig
	CircuitBreaker types.CircuitBreakerConfig
	HTTP           types.HTTPConfig
}

func Init() (*Config, error) {
	CFG = &Config{}
	ok := ReadConfig(CFG)
	if !ok {
		return nil, errors.New("[Init] Failed to read config file")
	}
	return CFG, nil
}

func ReadConfig(cfg *Config) bool {

	var configJson ConfigJson
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&configJson)

	cfg.Database.Master = configJson.Database.Master
	cfg.Database.Slave = configJson.Database.Slave

	cfg.Server.Port = configJson.Server.Port
	cfg.Server.GRPCPort = configJson.Server.GrpcPort

	return true
}
