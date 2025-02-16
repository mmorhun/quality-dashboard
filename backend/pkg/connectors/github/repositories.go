package github

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/go-github/v44/github"
)

var (
	GENERIC_ERROR_GIT_RESPONSE = errors.New("unable to parse repository. Please verify your token")
)

var (
	zeroDate = time.Time{}
)

func (g *Github) GetGithubRepositoryInformation(organization string, repository string) (*github.Repository, error) {
	fmt.Println("hola")
	fmt.Println(organization)
	fmt.Println(repository)
	fmt.Println("hola")

	repo, resp, err := g.client.Repositories.Get(context.Background(), organization, repository)
	if resp.StatusCode != 200 {
		return nil, GENERIC_ERROR_GIT_RESPONSE
	}
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (g *Github) GetRepositoryWorkflows(organization string, repository string) (*github.Workflows, error) {
	wk, resp, err := g.client.Actions.ListWorkflows(context.Background(), organization, repository, &github.ListOptions{})
	if resp.StatusCode != 200 {
		return nil, GENERIC_ERROR_GIT_RESPONSE
	}

	if err != nil {
		return nil, err
	}
	return wk, nil
}
