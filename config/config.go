package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ListenAddress       string   `json:"listen_address"`
	BackendServers      []string `json:"backend_servers"`
	HealthCheckInterval int      `json:"health_check_interval"`
	HealthCheckTimeout  int      `json:"health_check_timeout"`
}

func LoadConfig(configpath string) (*Config, error) {

	file, err := os.Open(configpath)

	if err != nil {
		fmt.Printf("Error opening file %s", err)
		return nil, err
	}

	defer file.Close()

	bytes, err1 := ioutil.ReadAll(file)

	if err1 != nil {
		fmt.Printf("Error reading file %s", err1)
		return nil, err1
	}

	var config Config

	err2 := json.Unmarshal(bytes, &config)

	if err2 != nil {
		log.Fatalf("Error unmarshalling json: %v", err2)
		return nil, err2
	}

	return &config, nil

}
