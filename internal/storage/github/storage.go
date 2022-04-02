package github

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v43/github"
	"github.com/shivamMg/ppds/tree"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

type Storage struct {
	spec   *spec
	client *github.Client
}

const gboxSuffix = ".gbox"

var (
	invalidName = fmt.Errorf("invalid name for secret")
	keyNotFound = errors.New("key not found")
)

func NewStorage(m map[string]interface{}) (*Storage, error) {
	s, err := wrapSpec(m)
	if err != nil {
		return nil, err
	}
	return &Storage{s, nil}, nil
}

func (s *Storage) Get(name string) (string, error) {
	name = name + gboxSuffix
	if err := s.checkRepoExist(); err != nil {
		return "", err
	}
	fc, err := s.getData(name)
	if err != nil {
		return "", err
	}
	if fc == nil {
		return "", keyNotFound
	}
	return fc.GetContent()
}

func (s *Storage) Set(name, data string) error {
	if strings.HasSuffix(name, "/") {
		return invalidName
	}

	ctx := context.Background()
	client := s.getClient()
	name = name + gboxSuffix

	if err := s.checkRepoExist(); err != nil {
		return err
	}

	commit := &github.RepositoryContentFileOptions{
		Content: []byte(data),
		Message: github.String(":inbox_tray: gbox: upload data"),
	}

	fc, err := s.getData(name)
	if err != nil {
		return err
	}

	// file not found
	if fc == nil {
		// create new file
		if _, _, err = client.Repositories.CreateFile(ctx, s.spec.owner, s.spec.repo, name, commit); err != nil {
			return err
		}
		return nil
	}

	// update file
	commit.SHA = fc.SHA
	if _, _, err = client.Repositories.UpdateFile(ctx, s.spec.owner, s.spec.repo, name, commit); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(name string) error {
	ctx := context.Background()
	client := s.getClient()
	name = name + gboxSuffix

	if err := s.checkRepoExist(); err != nil {
		return err
	}

	fc, err := s.getData(name)
	if err != nil {
		return err
	}
	if fc == nil {
		return keyNotFound
	}

	commit := &github.RepositoryContentFileOptions{
		Message: github.String(":outbox_tray: gbox: delete data"),
		SHA:     fc.SHA,
	}

	if _, _, err = client.Repositories.DeleteFile(ctx, s.spec.owner, s.spec.repo, name, commit); err != nil {
		return err
	}
	return nil
}

func (s *Storage) List() (tree.Node, error) {
	ctx := context.Background()
	client := s.getClient()

	if err := s.checkRepoExist(); err != nil {
		return nil, err
	}

	b, _, err := client.Repositories.GetBranch(ctx, s.spec.owner, s.spec.repo, s.spec.Branch, true)
	if err != nil {
		return nil, err
	}

	t, _, err := client.Git.GetTree(ctx, s.spec.owner, s.spec.repo, b.GetCommit().GetSHA(), true)
	if err != nil {
		return nil, err
	}

	return gitTreeToNode(t)
}

func (s *Storage) checkRepoExist() error {
	ctx := context.Background()
	client := s.getClient()
	if _, _, err := client.Repositories.Get(ctx, s.spec.owner, s.spec.repo); err != nil {
		return err
	}
	return nil
}

func (s *Storage) getClient() *github.Client {
	if s.client != nil {
		return s.client
	}

	var tc *http.Client
	if s.spec.AuthType == "https" {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: s.spec.Token},
		)
		tc = oauth2.NewClient(ctx, ts)
	}
	s.client = github.NewClient(tc)
	return s.client
}

func (s *Storage) getData(name string) (*github.RepositoryContent, error) {
	client := s.getClient()
	ctx := context.Background()
	fc, _, res, err := client.Repositories.GetContents(
		ctx,
		s.spec.owner,
		s.spec.repo,
		name,
		&github.RepositoryContentGetOptions{
			Ref: s.spec.Branch,
		})
	if err != nil {
		if res == nil {
			return nil, err
		}
		if res.StatusCode == http.StatusNotFound {
			return nil, nil
		}
	}
	if fc == nil {
		return nil, invalidName
	}
	return fc, nil
}
