package service

import (
	"time"

	"database/sql"
	"errors"

	"github.com/Yakshith15/blog-app/blog-service/internal/model"
	"github.com/Yakshith15/blog-app/blog-service/internal/repository"
	"github.com/google/uuid"
)

type BlogService struct {
	repo *repository.BlogRepository
}

func NewBlogService(repo *repository.BlogRepository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) CreateBlog(
	authorID uuid.UUID,
	title string,
	content string,
) (model.Blog, error) {

	blog := model.Blog{
		ID:        uuid.New(),
		AuthorID:  authorID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.Create(blog)
	return blog, err
}

func (s *BlogService) FindAll() ([]model.Blog, error) {
	return s.repo.FindAll()
}

func (s *BlogService) GetByID(id uuid.UUID) (model.Blog, error) {
	return s.repo.GetByID(id)
}

func (s *BlogService) UpdateBlog(
	id uuid.UUID,
	authorID uuid.UUID,
	title string,
	content string,
) (model.Blog, error) {
	blog, err := s.repo.UpdateBlog(id, authorID, title, content)
	if err == sql.ErrNoRows {
		return model.Blog{}, errors.New("not owner or blog not found")
	}
	return blog, err
}

func (s *BlogService) DeleteBlog(
	blogID uuid.UUID,
	requesterID uuid.UUID,
) error {
	err := s.repo.DeleteBlog(blogID, requesterID)
	if err == sql.ErrNoRows {
		return errors.New("not owner or blog not found")
	}
	return err
}
