package handler

import (
	"net/http"

	"github.com/Kenya-i/linu-quest/internal/domain"
	"github.com/gin-gonic/gin"
)

type StageHandler struct {
	stageUsecase domain.StageUsecase
}

func NewStageHandler(stageUsecase domain.StageUsecase) *StageHandler {
	return &StageHandler{stageUsecase: stageUsecase}
}

func (h *StageHandler) List(c *gin.Context) {
	stages, err := h.stageUsecase.GetStages(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stages)
}
