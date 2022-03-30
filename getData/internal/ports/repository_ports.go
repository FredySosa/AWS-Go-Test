package ports

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/getData/internal/core/domain"
)

type PostsRepositoryPort interface {
	GetPosts(ctx context.Context) ([]domain.Post, error)
}
