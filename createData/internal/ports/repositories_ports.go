package ports

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
)

//go:generate mockgen -destination=../mocks/repositories_ports_mock.go -package=mocks -source=repositories_ports.go

type PostsRepositoryPort interface {
	CreatePost(ctx context.Context, post domain.PostToSave) error
}
