package ports

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/getData/internal/core/domain"
)

type PostsServicePort interface {
	GetPosts(ctx context.Context) (domain.Response, error)
}
