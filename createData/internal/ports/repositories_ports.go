package ports

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
)

type PostsRepositoryPort interface {
	CreatePost(ctx context.Context, post domain.PostToSave) error
}
