package handler

import (
	"go-portfolio/server/api/response"
	"go-portfolio/server/lib/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- PROJECTS HANDLER ---

// GetProjects Get all projects
// @Summary Get All Projects
// @Description Take all projects data
// @Tags Projects
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse[[]models.Projects]
// @Router /projects [get]
func (h *Handler) GetProjects(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.LogRequest(ctx, logger.RequestData{
		Function:  "GetProjects",
		ProcessID: c.GetHeader("X-ProcessId"),
		IPAddress: c.ClientIP(),
		Request:   nil,
	})

	projects, err := h.project.GetAllProjects(ctx)
	if err != nil {
		res := response.ErrorResponse(http.StatusInternalServerError, "Failed to get projects")
		h.logger.LogResponse("99", res)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccessResponse(projects)
	h.logger.LogResponse("00", res)
	c.JSON(http.StatusOK, res)
}

// GetProjectDetail Get specific project
// @Summary Get Project Detail
// @Description Get specific project data by ID
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} response.BaseResponse[models.Projects]
// @Router /projects/{id} [get]
func (h *Handler) GetProjectDetail(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	project, err := h.project.GetProjectByID(ctx, id)
	if err != nil {
		res := response.ErrorResponse(http.StatusInternalServerError, err.Error())
		h.logger.LogResponse("99", res)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccessResponse(project)
	h.logger.LogResponse("00", res)
	c.JSON(http.StatusOK, res)
}
