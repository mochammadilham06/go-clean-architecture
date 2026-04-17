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

type ExperienceService struct {
	repo   contract.IExperienceRepository
	logger *logger.Logger
}

func NewExperienceService(repo contract.IExperienceRepository, logger *logger.Logger) *ExperienceService {
	return &ExperienceService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ExperienceService) GetAllExperience(ctx context.Context) ([]models.Experience, error) {

	res, err := s.repo.FindAll(ctx)

	if err != nil {
		s.logger.Error("[Func]: GetAllExperience : Failed to get all experience", zap.Error(err))
		return nil, response.InternalError
	}

	return res, nil

}

func (s *ExperienceService) GetExperienceByID(ctx context.Context, id string) (models.Experience, error) {

	if _, err := uuid.Parse(id); err != nil {
		s.logger.Warn("[Func] : GetExperienceByID :invalid input syntax for type uuid", zap.String("id", id))
		return models.Experience{}, response.ErrInvalidID
	}

	res, err := s.repo.FindByID(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Warn("[Func] : GetExperienceByID : Experience Not Found on list", zap.String("id", id))
			return models.Experience{}, response.ErrProjectNotFound
		}

		s.logger.Error("[Func]: GetExperienceByID : No Data Found table", zap.Error(err))
		return models.Experience{}, response.ErrProjectNotFound
	}

	return res, nil
}
