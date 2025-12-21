package service

import (
	"database/sql"
	"errors"
	"time"
	"github.com/google/uuid"
	"github.com/Yakshith15/blog-app/comment-service/internal/model"
	"github.com/Yakshith15/blog-app/comment-service/internal/repository"
)

type CommentService struct {
	commentRepository *repository.CommentRepository
}

func NewCommentService(commentRepository *repository.CommentRepository) *CommentService {
	return &CommentService{commentRepository: commentRepository}
}

func (s *CommentService) CreateComment(blogID uuid.UUID, authorID uuid.UUID, content string) (model.Comment, error) {
	comment := model.Comment{
		ID: uuid.New(),
		BlogId: blogID,
		AuthorId: authorID,
		Content: content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.commentRepository.Create(comment)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (s *CommentService) GetCommentsByBlogID(blogID uuid.UUID) ([]model.Comment, error) {
	return s.commentRepository.GetCommentsByBlogID(blogID)
}

func (s *CommentService) GetCommentByID(id uuid.UUID) (*model.Comment, error) {
	return s.commentRepository.GetCommentByID(id)
}

func (s *CommentService) UpdateComment(id uuid.UUID, authorID uuid.UUID, blogID uuid.UUID, content string) (model.Comment, error) {
	comment := model.Comment{
		ID: id,
		BlogId: blogID,
		AuthorId: authorID,
		Content: content,
		UpdatedAt: time.Now(),
	}
	err := s.commentRepository.Update(comment)
	if err == sql.ErrNoRows {
		return model.Comment{}, errors.New("not owner or comment not found")
	}
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (s *CommentService) DeleteComment(id uuid.UUID, authorID uuid.UUID, blogID uuid.UUID) error {
	err := s.commentRepository.Delete(id, authorID, blogID)
	if err == sql.ErrNoRows {
		return errors.New("not owner or comment not found")
	}
	if err != nil {
		return err
	}
	return nil
}