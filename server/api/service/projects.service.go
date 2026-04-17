package service

import (
	"context"
	"database/sql"
	"errors"
	"go-portfolio/server/api/contract"
	"go-portfolio/server/api/models"
	"go-portfolio/server/api/response"
	"go-portfolio/server/lib/logger"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ProjectService struct {
	repo   contract.IProjectRepository
	logger *logger.Logger
}

func NewProjectService(repo contract.IProjectRepository, log *logger.Logger) *ProjectService {
	return &ProjectService{
		repo:   repo,
		logger: log,
	}
}

func (s *ProjectService) GetAllProjects(ctx context.Context) ([]models.Projects, error) {
	projects, err := s.repo.FindAll(ctx)
	if err != nil {

		s.logger.Error("failed to fetch all projects", zap.Error(err))
		return nil, err
	}
	return projects, nil
}

func (s *ProjectService) GetProjectByID(ctx context.Context, id string) (models.Projects, error) {

	if _, err := uuid.Parse(id); err != nil {
		s.logger.Warn("invalid input syntax for type uuid", zap.String("id", id))
		return models.Projects{}, response.ErrInvalidID
	}

	project, err := s.repo.FindByID(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Warn("Project Not Found on list", zap.String("id", id))
			return models.Projects{}, response.ErrProjectNotFound
		}

		s.logger.Error("No Data Found table", zap.Error(err))
		return models.Projects{}, response.ErrProjectNotFound
	}
	return project, nil
}
