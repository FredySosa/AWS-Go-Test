package services

import (
	"context"
	"log"

	"github.com/FredySosa/AWS-Go-Test/createData/internal/core/domain"
	"github.com/FredySosa/AWS-Go-Test/createData/internal/ports"
	"github.com/google/uuid"
)

type PostsService struct {
	PostsRepository ports.PostsRepositoryPort
}

func NewPostsService(pr ports.PostsRepositoryPort) PostsService {
	return PostsService{
		PostsRepository: pr,
	}
}

func (ps PostsService) CreatePort(ctx context.Context, request domain.CreationRequest) (domain.Response, error) {
	if request.URLToImage == "" || request.Text == "" {
		return domain.Response{}, domain.ValidationError
	}

	toCreate := domain.PostToSave{
		ID:         uuid.New().String(),
		URLToImage: request.URLToImage,
		Text:       request.Text,
	}

	if err := ps.PostsRepository.CreatePost(ctx, toCreate); err != nil {
		log.Println(err)
		return domain.Response{}, domain.UnknownErr
	}

	return domain.Response{
		ID: toCreate.ID,
	}, nil
}
