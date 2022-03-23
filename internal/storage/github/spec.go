package github

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
)

type spec struct {
	URL      string `yaml:"url"`
	AuthType string `yaml:"authType"`
	Token    string `yaml:"token"`
}

func wrapSpec(m map[string]interface{}) (*spec, error) {
	errTpl := "github spec invalid: %s"
	b, err := yaml.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf(errTpl, err)
	}

	s := new(spec)
	if err = yaml.Unmarshal(b, s); err != nil {
		return nil, fmt.Errorf(errTpl, err)
	}

	// validate spec
	if err = s.validate(); err != nil {
		return nil, fmt.Errorf(errTpl, err)
	}
	return s, nil
}

func (s *spec) validate() error {
	if len(s.URL) == 0 {
		return errors.New("url required")
	}
	switch s.AuthType {
	case "":
		return nil
	case "https":
		if len(s.Token) == 0 {
			return errors.New("token required")
		}
	default:
		return fmt.Errorf("unknown auth type: %s", s.AuthType)
	}
	return nil
}
