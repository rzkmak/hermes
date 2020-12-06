package reader

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type SshConfigs struct {
	Config map[string]SshConfig
}

type SshConfig struct {
	Hostname string `yaml:"hostname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Key      string `yaml:"key"`
}

func ReadYaml() (SshConfigs, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := fmt.Sprintf("%v/.hermes", home)
	files, err := ioutil.ReadDir(configDir)
	if err != nil {
		return SshConfigs{}, err
	}
	configs := map[string]SshConfig{}
	for _, f := range files {
		if !strings.Contains(f.Name(), "yaml") && !strings.Contains(f.Name(), "yml") {
			continue
		}
		yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", configDir, f.Name()))
		if err != nil {
			return SshConfigs{}, err
		}
		var config []SshConfig
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			return SshConfigs{}, err
		}
		for i := 0; i < len(config); i++ {
			if config[i].Password != "" {
				password, err := base64.StdEncoding.DecodeString(config[i].Password)
				if err != nil {
					log.Fatalln("decode password error: ", err)
				}
				config[i].Password = string(password)
			}
			if config[i].Key != "" {
				key, err := base64.StdEncoding.DecodeString(config[i].Key)
				if err != nil {
					log.Fatalln("decode key error: ", err)
				}
				config[i].Key = string(key)
			}
			configs[config[i].Hostname] = config[i]
		}
	}
	if len(configs) == 0 {
		return SshConfigs{}, errors.New("no configuration yaml found")
	}
	return SshConfigs{Config: configs}, nil
}
