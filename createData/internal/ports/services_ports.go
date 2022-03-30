package ports

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
)

type PostsServicePort interface {
	CreatePost(ctx context.Context, request domain.CreationRequest) (domain.Response, error)
}
