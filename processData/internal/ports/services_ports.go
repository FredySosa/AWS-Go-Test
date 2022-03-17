package ports

import (
	"context"

	"github.com/FredySosa/AWS-Go-Test/processData/internal/core/domain"
)

type SNSServicePort interface {
	PublishMessage(ctx context.Context, posts []domain.Post) error
}
