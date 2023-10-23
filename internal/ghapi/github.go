package ghapi

import (
	"context"
	"errors"
	"time"

	"github.com/amir/graphql-c1/pkg/auth"
	"github.com/google/go-github/v56/github"

	"github.com/amir/graphql-c1/internal/model"
)

type GhAPI struct {
	client *github.Client
}

func New() *GhAPI {
	return &GhAPI{
		client: github.NewClient(nil),
	}
}

func (s *GhAPI) PublicGists(ctx context.Context, since time.Time, page int, perPage int) ([]*model.Gist, error) {
	gists, _, err := s.client.Gists.ListAll(ctx, &github.GistListOptions{
		Since: since,
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	})
	if err != nil {
		return nil, err
	}

	return convertAPIToGists(gists), nil
}

func (s *GhAPI) GistByID(ctx context.Context, id string) (*model.Gist, error) {
	gist, _, err := s.client.Gists.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertAPIToGist(gist), nil
}

func (s *GhAPI) GistComments(ctx context.Context, gistID string, page, perPage int) ([]*model.Comment, error) {
	comments, _, err := s.client.Gists.ListComments(ctx, gistID, &github.ListOptions{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		return nil, err
	}

	return convertAPIToComments(comments), nil
}

func (s *GhAPI) GistCommits(ctx context.Context, gistID string, page, perPage int) ([]*model.Commit, error) {
	commits, _, err := s.client.Gists.ListCommits(ctx, gistID, &github.ListOptions{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		return nil, err
	}

	return convertAPIToCommits(commits), nil
}

var (
	ErrUnauthorized = errors.New("authorization token not provided")
)

func (s *GhAPI) CreateGist(ctx context.Context, in *model.CreateGistRequest) (*model.Gist, error) {
	files := make(map[github.GistFilename]github.GistFile)

	for _, file := range in.Files {
		key := github.GistFilename(file.FileName)

		files[key] = github.GistFile{
			Filename: github.String(file.FileName),
			Content:  github.String(file.Content),
		}
	}

	token, ok := auth.FromContext(ctx)
	if !ok {
		return nil, ErrUnauthorized
	}

	client := s.client.WithAuthToken(token)

	gist, _, err := client.Gists.Create(ctx, &github.Gist{
		Description: github.String(in.Description),
		Public:      github.Bool(in.Public),
		Files:       files,
	})
	if err != nil {
		return nil, err
	}

	return convertAPIToGist(gist), nil
}

func (s *GhAPI) UpdateGist(ctx context.Context, id string, in *model.CreateGistRequest) (*model.Gist, error) {
	files := make(map[github.GistFilename]github.GistFile)

	for _, file := range in.Files {
		key := github.GistFilename(file.FileName)

		files[key] = github.GistFile{
			Filename: github.String(file.FileName),
			Content:  github.String(file.Content),
		}
	}

	token, ok := auth.FromContext(ctx)
	if !ok {
		return nil, ErrUnauthorized
	}

	client := s.client.WithAuthToken(token)

	gist, _, err := client.Gists.Edit(ctx, id, &github.Gist{
		Description: github.String(in.Description),
		Public:      github.Bool(in.Public),
		Files:       files,
	})
	if err != nil {
		return nil, err
	}

	return convertAPIToGist(gist), nil
}
