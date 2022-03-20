package storage

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	defaultConfigPath = ".config/gbox"
	defaultConfigFile = "config.yaml"
)

type config struct {
	CurrentStorage string `yaml:"currentStorage"`
	Storages       []*struct {
		Name string                 `yaml:"name"`
		Kind string                 `yaml:"kind"`
		Spec map[string]interface{} `yaml:"spec"`
	} `yaml:"storages"`
}

var _config *config

func GetConfig() (*config, error) {
	if cfg, err := loadConfig(); err == nil {
		return cfg, nil
	}
	return &config{}, nil
}

func getConfigPath() string {
	cp := os.Getenv("GBOX_CONFIG_PATH")
	if len(cp) == 0 {
		cp = defaultConfigPath
	}
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, cp)
}

func getConfigFile() string {
	cf := os.Getenv("GBOX_CONFIG_FILE")
	if len(cf) == 0 {
		cf = defaultConfigFile
	}
	return cf
}

func loadConfig() (*config, error) {
	if _config != nil {
		return _config, nil
	}

	configDir, err := filepath.Abs(filepath.Join(getConfigPath(), getConfigFile()))
	if err != nil {
		return nil, err
	}

	f, err := os.Open(configDir)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	cfg, err := parseConfig(b)
	if err != nil {
		return nil, err
	}

	err = validateConfig(cfg)
	if err != nil {
		return nil, err
	}

	_config = cfg
	return _config, nil
}

func parseConfig(b []byte) (*config, error) {
	cfg := new(config)
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func validateConfig(cfg *config) error {
	for i, s := range cfg.Storages {
		if len(s.Name) == 0 {
			return fmt.Errorf("config validation failed: storages[%d].name: required", i)
		}
		if len(s.Kind) == 0 {
			return fmt.Errorf("config validation failed: storages[%d].kind: required", i)
		}
	}
	return nil
}

func (cfg *config) Save() error {
	cp := getConfigPath()
	if err := os.MkdirAll(cp, os.ModePerm); err != nil {
		return err
	}

	cf := getConfigFile()
	configDir, err := filepath.Abs(filepath.Join(cp, cf))
	if err != nil {
		return err
	}
	d, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(configDir, d, 0); err != nil {
		return err
	}
	return nil
}
