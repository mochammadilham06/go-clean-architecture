package repository

import (
	"context"
	"database/sql"
	"go-portfolio/server/api/models"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) FindAll(ctx context.Context) ([]models.Projects, error) {
	query := `SELECT id, title, category, year, image_url, description, created_at 
              FROM public.projects 
              ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Projects
	for rows.Next() {
		var p models.Projects
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Category,
			&p.Year,
			&p.ImageUrl,
			&p.Description,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *ProjectRepository) FindByID(ctx context.Context, id string) (models.Projects, error) {
	query := `SELECT id, title, category, year, image_url, description, created_at 
              FROM public.projects 
              WHERE id = $1`

	var p models.Projects
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&p.ID,
		&p.Title,
		&p.Category,
		&p.Year,
		&p.ImageUrl,
		&p.Description,
		&p.CreatedAt,
	)

	if err != nil {
		return models.Projects{}, err
	}

	return p, nil
}
