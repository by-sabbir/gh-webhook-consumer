package events

import (
	"context"

	hook "github.com/by-sabbir/gh-webhook-consumer/internal/hooks"
)

type Service struct {
	Store Store
}

// NewService - returns to a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// Store - this interface defines all the methods to operate
type Store interface {
	GetIssues(context.Context, hook.GitIssues) (string, error)
}

func (s *Service) GetIssues(ctx context.Context, gi hook.GitIssues) (string, error) {
	insertedId, err := s.Store.GetIssues(ctx, gi)

	if err != nil {
		return "", err
	}

	return insertedId, nil
}
