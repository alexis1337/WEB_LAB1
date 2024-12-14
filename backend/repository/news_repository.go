package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"news_app/models"
)

var ErrNewsNotFound = errors.New("новость не найдена")

type NewsRepository interface {
	GetAll(ctx context.Context) ([]models.News, error)
	GetNewsByID(ctx context.Context, id int) (models.News, error)
	Create(ctx context.Context, news models.News) (int, error)
	Update(ctx context.Context, id int, news models.News) error
	Delete(ctx context.Context, id int) error
}

type newsRepository struct {
	db *sql.DB
}

func NewNewsRepository(db *sql.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) GetAll(ctx context.Context) ([]models.News, error) {
	query := `
		SELECT 
			id, 
			title, 
			author, 
			content 
		FROM news`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	var newsList []models.News
	for rows.Next() {
		var news models.News
		if err := rows.Scan(&news.ID, &news.Title, &news.Author, &news.Content); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}
		newsList = append(newsList, news)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("ошибка при итерации по строкам: %w", rows.Err())
	}

	return newsList, nil
}

func (r *newsRepository) GetNewsByID(ctx context.Context, id int) (models.News, error) {
	query := `
		SELECT 
			id, 
			title, 
			author, 
			content 
		FROM news 
		WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var news models.News
	err := row.Scan(&news.ID, &news.Title, &news.Author, &news.Content)
	if errors.Is(err, sql.ErrNoRows) {
		return models.News{}, ErrNewsNotFound
	}
	if err != nil {
		return models.News{}, fmt.Errorf("ошибка при получении новости по ID: %w", err)
	}

	return news, nil
}

func (r *newsRepository) Create(ctx context.Context, news models.News) (int, error) {
	query := `
		INSERT INTO news (title, author, content) 
		VALUES ($1, $2, $3) 
		RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, query, news.Title, news.Author, news.Content).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания новости: %w", err)
	}
	return id, nil
}

func (r *newsRepository) Update(ctx context.Context, id int, news models.News) error {
	query := `
		UPDATE news 
		SET 
			title = $1, 
			author = $2, 
			content = $3 
		WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, news.Title, news.Author, news.Content, id)
	if err != nil {
		return fmt.Errorf("ошибка обновления новости по ID %d: %w", id, err)
	}
	return nil
}

func (r *newsRepository) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM news 
		WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления новости по ID %d: %w", id, err)
	}
	return nil
}
