package storage

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/logger"
	"github.com/zezaeoh/gbox/internal/storage"
	"github.com/zezaeoh/gbox/internal/storage/github"
	"net/url"
	"regexp"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add storage config",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()

		ans := struct {
			Name string
			Kind string
			Spec map[string]interface{}
		}{}

		q := []*survey.Question{{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Name of storage",
				Help:    "Name to use identify storage",
			},
			Validate: func(val interface{}) error {
				name, ok := val.(string)
				if !ok {
					return errors.New("name must be string")
				}
				match, _ := regexp.MatchString("^[a-z0-9]([-a-z0-9]*[a-z0-9])?$", name)
				if !match {
					return errors.New("not matched name with regex `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`")
				}
				return nil
			},
		}, {
			Name: "kind",
			Prompt: &survey.Select{
				Message: "Kind of storage",
				Help:    "What kind of storage do you want to create?",
				Options: []string{"github"},
			},
		}}
		if err := survey.Ask(q, &ans); err != nil {
			log.Errorf("Fail to get answer: %s", err)
			return
		}

		switch ans.Kind {
		case "github":
			s := make(map[string]interface{}, 3)
			q := []*survey.Question{{
				Name: "url",
				Prompt: &survey.Input{
					Message: "Repository URL",
					Help:    "URL of your public/private github repository",
				},
				Validate: func(val interface{}) error {
					uri, ok := val.(string)
					if !ok {
						return errors.New("URL must be string")
					}
					if _, err := url.ParseRequestURI(uri); err != nil {
						return err
					}
					return nil
				},
			}, {
				Name: "branch",
				Prompt: &survey.Input{
					Message: "Repository Branch",
					Help:    "Branch of your repository to use",
				},
				Validate: survey.Required,
			}, {
				Name: "authType",
				Prompt: &survey.Select{
					Message: "Authentication method",
					Help:    "Select authentication method of your repository. Choose none if repository is public",
					Options: []string{"none", "https"},
				},
			}}
			if err := survey.Ask(q, &s); err != nil {
				log.Errorf("Fail to get spec: %s", err)
				return
			}

			// extract value from option answer type
			s["authType"] = s["authType"].(survey.OptionAnswer).Value

			// authType check
			if s["authType"] == "none" {
				delete(s, "authType")
			} else if s["authType"] == "https" {
				token := ""
				if err := survey.AskOne(&survey.Input{
					Message: "Personal access token",
					Help:    "Please see https://github.blog/2020-12-15-token-authentication-requirements-for-git-operations/ for more information.",
				}, &token, survey.WithValidator(survey.Required)); err != nil {
					log.Errorf("Fail to get token: %s", err)
					return
				}
				s["token"] = token
			}

			// validate storage spec
			if _, err := github.NewStorage(s); err != nil {
				log.Errorf("Fail to validate github storage: %s", err)
				return
			}

			// save configuration
			cfg, err := storage.GetConfig()
			if err != nil {
				log.Errorf("Fail to get config: %s", err)
				return
			}

			// add storage to config file and save
			cfg.AddStorage(ans.Name, ans.Kind, s)
			if err := cfg.Save(); err != nil {
				log.Errorf("Fail to save config: %s", err)
				return
			}
		default:
			log.Errorf("Unkwon kind of storage: %s", ans.Kind)
			return
		}
		log.Infof("Successfully add %s storage config: %s", ans.Kind, ans.Name)
	},
}
