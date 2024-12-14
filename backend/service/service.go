package service

import (
	"context"
	"news_app/models"
	"news_app/repository"
)

type NewsService interface {
	GetAllNews(ctx context.Context) ([]models.News, error)
	GetNewsByID(ctx context.Context, id int) (models.News, error)
	CreateNews(ctx context.Context, news models.News) (models.News, error)
	UpdateNews(ctx context.Context, id int, news models.News) (models.News, error)
	DeleteNews(ctx context.Context, id int) error
}

type newsService struct {
	repo repository.NewsRepository
}

func NewNewsService(repo repository.NewsRepository) NewsService {
	return &newsService{repo: repo}
}

func (s *newsService) GetAllNews(ctx context.Context) ([]models.News, error) {
	return s.repo.GetAll(ctx)
}

func (s *newsService) GetNewsByID(ctx context.Context, id int) (models.News, error) {
	return s.repo.GetNewsByID(ctx, id)
}

func (s *newsService) CreateNews(ctx context.Context, news models.News) (models.News, error) {
	id, err := s.repo.Create(ctx, news)
	if err != nil {
		return models.News{}, err
	}
	news.ID = id
	return news, nil
}

func (s *newsService) UpdateNews(ctx context.Context, id int, news models.News) (models.News, error) {
	err := s.repo.Update(ctx, id, news)
	if err != nil {
		return models.News{}, err
	}
	news.ID = id
	return news, nil
}

func (s *newsService) DeleteNews(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
