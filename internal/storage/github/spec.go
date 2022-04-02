package github

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type spec struct {
	URL      string `yaml:"url"`
	Branch   string `yaml:"branch"`
	AuthType string `yaml:"authType"`
	Token    string `yaml:"token"`

	owner string
	repo  string
}

const githubDomain = "github.com"

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
	if len(s.Branch) == 0 {
		return errors.New("branch required")
	}

	// extract owner, repo from url
	invalidUrlErr := fmt.Errorf("invalid github url: %s", s.URL)
	idx := strings.Index(s.URL, githubDomain)
	if idx == -1 {
		return invalidUrlErr
	}
	ss := strings.Split(s.URL[idx+len(githubDomain)+1:len(s.URL)], "/")
	if len(ss) != 2 {
		return invalidUrlErr
	}
	s.owner = ss[0]
	s.repo = ss[1]
	if strings.HasSuffix(s.repo, ".git") {
		s.repo = s.repo[:len(s.repo)-4]
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
