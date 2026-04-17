package handler

import (
	"go-portfolio/server/api/response"
	"go-portfolio/server/lib/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- EXPERIENCE HANDLER ---

// GetExperiences Get All Experience
// @Summary Get All Experience
// @Description Take all experience data
// @Tags Experience
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse[[]models.Experience]
// @Router /experience [get]
func (h *Handler) GetExperiences(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.LogRequest(ctx, logger.RequestData{
		Function:  "GetExperiences",
		ProcessID: c.GetHeader("X-ProcessId"),
		IPAddress: c.ClientIP(),
		Request:   nil,
	})

	exp, err := h.experience.GetAllExperience(ctx)

	if err != nil {
		res := response.ErrorResponse(http.StatusInternalServerError, err.Error())
		h.logger.LogResponse("99", res)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccessResponse(exp)
	h.logger.LogResponse("00", res)
	c.JSON(http.StatusOK, res)

}

// GetExperiencesDetail Get Detail All Experience
// @Summary Get Detail Experience
// @Description Get deail experience data
// @Tags Experience
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse[models.Experience]
// @Router /experience/{id} [get]
func (h *Handler) GetExperiencesDetail(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	h.logger.LogRequest(ctx, logger.RequestData{
		Function:  "GetExperiencesDetail",
		ProcessID: "123",
		IPAddress: c.ClientIP(),
		Request:   id,
	})

	exp, err := h.experience.GetExperienceByID(ctx, id)

	if err != nil {
		res := response.ErrorResponse(http.StatusInternalServerError, err.Error())
		h.logger.LogResponse("99", res)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccessResponse(exp)
	h.logger.LogResponse("00", res)

	c.JSON(http.StatusOK, res)

}
