package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/Yakshith15/blog-app/blog-service/internal/model"
	"time"
)

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) Create(blog model.Blog) error {

	query := `
		INSERT INTO blogs (id, author_id, title, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		query,
		blog.ID,
		blog.AuthorID,
		blog.Title,
		blog.Content,
		blog.CreatedAt,
		blog.UpdatedAt,
	)

	return err
}

func (r *BlogRepository) FindAll() ([]model.Blog, error) {

	rows, err := r.db.Query(`
		SELECT id, author_id, title, content, created_at
		FROM blogs
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.Blog

	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(
			&blog.ID,
			&blog.AuthorID,
			&blog.Title,
			&blog.Content,
			&blog.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}


func (r *BlogRepository) GetByID(id uuid.UUID) (model.Blog, error) {

	row := r.db.QueryRow(`
		SELECT id, author_id, title, content, created_at
		FROM blogs
		WHERE id = $1
	`, id)

	var blog model.Blog
	err := row.Scan(
		&blog.ID,
		&blog.AuthorID,
		&blog.Title,
		&blog.Content,
		&blog.CreatedAt,
	)
	return blog, err
}


func (r *BlogRepository) UpdateBlog(
	id uuid.UUID,
	authorID uuid.UUID,
	title string,
	content string,
) (model.Blog, error) {

	query := `
		UPDATE blogs
		SET title = $1,
		    content = $2,
		    updated_at = $3
		WHERE id = $4
		  AND author_id = $5
		RETURNING id, author_id, title, content, created_at, updated_at
	`

	var blog model.Blog

	err := r.db.QueryRow(
		query,
		title,
		content,
		time.Now(),
		id,
		authorID,
	).Scan(
		&blog.ID,
		&blog.AuthorID,
		&blog.Title,
		&blog.Content,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)

	return blog, err
}

func (r *BlogRepository) DeleteBlog(
	id uuid.UUID,
	authorID uuid.UUID,
) error {

	query := `
		DELETE FROM blogs
		WHERE id = $1
		  AND author_id = $2
	`

	result, err := r.db.Exec(query, id, authorID)
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

func (r *BlogRepository) ExistsByID(id uuid.UUID) (bool, error) {

	query := `
		SELECT 1
		FROM blogs
		WHERE id = $1
		LIMIT 1
	`

	var dummy int
	err := r.db.QueryRow(query, id).Scan(&dummy)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

