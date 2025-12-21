package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/Yakshith15/blog-app/comment-service/internal/model"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment model.Comment) error {
	query := `
		INSERT INTO comments (id, blogId, authorId, content, createdAt, updatedAt)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(
		query,
		comment.ID,
		comment.BlogId,
		comment.AuthorId,
		comment.Content,
		comment.CreatedAt,
		comment.UpdatedAt,
	)
	return err
}

func (r *CommentRepository) GetCommentsByBlogID(blogID uuid.UUID) ([]model.Comment, error) {
	query := `
		SELECT id, blogId, authorId, content, createdAt, updatedAt
		FROM comments
		WHERE blogId = $1
		ORDER BY createdAt ASC
	`
	rows, err := r.db.Query(query, blogID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.BlogId,
			&comment.AuthorId,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *CommentRepository) GetCommentByID(id uuid.UUID) (*model.Comment, error) {
	query := `
		SELECT id, blogId, authorId, content, createdAt, updatedAt
		FROM comments
		WHERE id = $1
	`
	var comment model.Comment
	err := r.db.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.BlogId,
		&comment.AuthorId,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *CommentRepository) Update(comment model.Comment) error {
	query := `
		UPDATE comments
		SET content = $1, updatedAt = $2
		WHERE id = $3 AND authorId = $4 AND blogId = $5
	`
	result, err := r.db.Exec(
		query,
		comment.Content,
		comment.UpdatedAt,
		comment.ID,
		comment.AuthorId,
		comment.BlogId,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *CommentRepository) Delete(id uuid.UUID, authorID uuid.UUID, blogID uuid.UUID) error {
	query := `
		DELETE FROM comments
		WHERE id = $1 AND authorId = $2 AND blogId = $3
	`
	result, err := r.db.Exec(query, id, authorID, blogID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
