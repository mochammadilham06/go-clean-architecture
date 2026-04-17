package contract

import (
	"context"
	"go-portfolio/server/api/models"
)

type IProjectService interface {
	GetAllProjects(ctx context.Context) ([]models.Projects, error)
	GetProjectByID(ctx context.Context, id string) (models.Projects, error)
}

type IProjectRepository interface {
	FindAll(ctx context.Context) ([]models.Projects, error)
	FindByID(ctx context.Context, id string) (models.Projects, error)
}
