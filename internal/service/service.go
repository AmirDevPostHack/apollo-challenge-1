package service

import (
	"context"
	"time"

	"github.com/amir/graphql-c1/internal/ghapi"
	"github.com/amir/graphql-c1/internal/model"
)

type Service struct {
	ghAPI *ghapi.GhAPI
}

func New() *Service {
	return &Service{
		ghAPI: ghapi.New(),
	}
}

func (s *Service) PublicGists(ctx context.Context, since time.Time, page, perPage int) ([]*model.Gist, error) {
	return s.ghAPI.PublicGists(ctx, since, page, perPage)
}

func (s *Service) GistByID(ctx context.Context, id string) (*model.Gist, error) {
	return s.ghAPI.GistByID(ctx, id)
}

func (s *Service) GistComments(ctx context.Context, gistID string, page, perPage int) ([]*model.Comment, error) {
	return s.ghAPI.GistComments(ctx, gistID, page, perPage)
}

func (s *Service) GistCommits(ctx context.Context, gistID string, page, perPage int) ([]*model.Commit, error) {
	return s.ghAPI.GistCommits(ctx, gistID, page, perPage)
}

func (s *Service) GistCreate(ctx context.Context, req model.CreateGistRequest) (*model.Gist, error) {
	return s.ghAPI.CreateGist(ctx, &req)
}

func (s *Service) GistUpdate(ctx context.Context, id string, req model.CreateGistRequest) (*model.Gist, error) {
	return s.ghAPI.UpdateGist(ctx, id, &req)
}
