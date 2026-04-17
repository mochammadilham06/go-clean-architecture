package contract

import (
	"context"
	"go-portfolio/server/api/models"
)

type IExperienceService interface {
	GetAllExperience(ctx context.Context) ([]models.Experience, error)
	GetExperienceByID(ctx context.Context, id string) (models.Experience, error)
}

type IExperienceRepository interface {
	FindAll(ctx context.Context) ([]models.Experience, error)
	FindByID(ctx context.Context, id string) (models.Experience, error)
}
