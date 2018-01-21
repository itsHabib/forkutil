package repo

import (
	"fmt"
	"path/filepath"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

// GHRepo represents the repository given by a user
type GHRepo struct {
	RepoDir string
	owner   string
	project string
	repo    *git.Repository
}

// NewGHRepo instantiates a new GHRepo struct
func NewGHRepo(repository string) (*GHRepo, error) {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return nil, fmt.Errorf("repository must be in format owner/project")
	}
	return &GHRepo{
		owner:   values[0],
		project: values[1],
	}, nil
}

// RepositoryURL returns the repository URL of a given repository
func (g *GHRepo) RepositoryURL() string {
	return fmt.Sprintf("https://github.com/%s/%s.git", g.owner, g.project)
}

// Clone clones repository onto user machine at the destination
func (g *GHRepo) Clone(dest string) error {
	fullPath := filepath.Join(dest, fmt.Sprintf("%s-%s", g.owner, g.project))
	repo, err := git.PlainClone(fullPath, false, &git.CloneOptions{
		URL: g.RepositoryURL(),
	})
	if err != nil {
		return err
	}
	g.repo = repo
	g.RepoDir = fullPath
	return err
}
