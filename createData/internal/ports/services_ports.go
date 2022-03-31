package ports

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
)

//go:generate mockgen -destination=../mocks/services_ports_mock.go -package=mocks -source=services_ports.go

type PostsServicePort interface {
	CreatePost(ctx context.Context, request domain.CreationRequest) (domain.Response, error)
}
