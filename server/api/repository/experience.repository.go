package repository

import (
	"context"
	"database/sql"
	"go-portfolio/server/api/models"
)

type ExperienceRepository struct {
	db *sql.DB
}

func NewExperienceRepository(db *sql.DB) *ExperienceRepository {
	return &ExperienceRepository{
		db: db,
	}
}

func (r *ExperienceRepository) FindAll(ctx context.Context) ([]models.Experience, error) {
	query := `SELECT ex."role", ex.description, ex.company, ex.start_date, ex.end_date, 
			  ex.is_active, ex.tech_stack, ex.link , ex.created_at , ex.updated_at 
              FROM public.experience ex 
              ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var works []models.Experience

	for rows.Next() {
		var w models.Experience
		err := rows.Scan(
			// &w.Id,
			&w.Role,
			&w.Description,
			&w.Company,
			&w.StartDate,
			&w.EndDate,
			&w.IsActive,
			&w.TechStack,
			&w.Link,
			&w.CreatedAt,
			&w.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		works = append(works, w)
	}

	return works, nil

}

func (r *ExperienceRepository) FindByID(ctx context.Context, id string) (models.Experience, error) {

	query := `SELECT ex."role", ex.description, ex.company, ex.start_date, ex.end_date, 
		  ex.is_active, ex.tech_stack, ex.link , ex.created_at , ex.updated_at 
          FROM public.experience ex  
          WHERE id = $1`

	var w models.Experience
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&w.Role,
		&w.Description,
		&w.Company,
		&w.StartDate,
		&w.EndDate,
		&w.IsActive,
		&w.TechStack,
		&w.Link,
		&w.CreatedAt,
		&w.UpdatedAt,
	)

	if err != nil {
		return models.Experience{}, err
	}

	return w, nil
}
