package services

import (
	"context"
	"log"

	"github.com/FredySosa/AWS-Go-Test/getData/internal/core/domain"
	"github.com/FredySosa/AWS-Go-Test/getData/internal/ports"
)

type PostsService struct {
	PostsRepository ports.PostsRepositoryPort
}

func NewPostsService(pr ports.PostsRepositoryPort) PostsService {
	return PostsService{
		PostsRepository: pr,
	}
}

func (ps PostsService) GetPosts(ctx context.Context) (domain.Response, error) {
	posts, err := ps.PostsRepository.GetPosts(ctx)
	if err != nil {
		log.Println(err)

		return domain.Response{}, domain.UnknownErr
	}

	return domain.Response{
		Posts: posts,
	}, nil
}
