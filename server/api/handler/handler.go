package handler

import (
	"go-portfolio/server/api/contract"
	"go-portfolio/server/lib/logger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	project    contract.IProjectService
	experience contract.IExperienceService
	logger     *logger.Logger
}

// Handler Contructor
func NewHandler(
	project contract.IProjectService,
	experience contract.IExperienceService,
	logger *logger.Logger,
) *Handler {
	return &Handler{
		project:    project,
		experience: experience,
		logger:     logger,
	}
}

func (h *Handler) Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Route Project (logic from project.handler.go)
		v1.GET("/projects", h.GetProjects)
		v1.GET("/projects/:id", h.GetProjectDetail)

		// Route Experience (logic from experience.handler.go)
		v1.GET("/experience", h.GetExperiences)
		v1.GET("/experience/:id", h.GetExperiences)

	}
}
