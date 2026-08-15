package handler

import (
	"net/http"

	"github.com/Kenya-i/linu-quest/internal/domain"
	"github.com/gin-gonic/gin"
)

type questionResponse struct {
	ID             string   `json:"id"`
	StageID        string   `json:"stage_id"`
	QuestionNumber int      `json:"question_number"`
	QuestionText   string   `json:"question_text"`
	Choices        []string `json:"choices"`
}

type QuestionHandler struct {
	questionUsecase domain.QuestionUsecase
}

func NewQuestionHandler(questionUsecase domain.QuestionUsecase) *QuestionHandler {
	return &QuestionHandler{questionUsecase: questionUsecase}
}

func (h *QuestionHandler) ListByStage(c *gin.Context) {
	stageID := c.Param("id")

	questions, err := h.questionUsecase.GetQuestionsByStageID(c.Request.Context(), stageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]questionResponse, len(questions))
	for i, q := range questions {
		responses[i] = questionResponse{
			ID:             q.ID,
			StageID:        q.StageID,
			QuestionNumber: q.QuestionNumber,
			QuestionText:   q.QuestionText,
			Choices:        q.Choices,
		}
	}

	c.JSON(http.StatusOK, responses)
}

type checkAnswerRequest struct {
	SelectedPosition int `json:"selected_position"`
}

func (h *QuestionHandler) CheckAnswer(c *gin.Context) {
	questionID := c.Param("id")

	var req checkAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	correct, err := h.questionUsecase.CheckAnswer(c.Request.Context(), questionID, req.SelectedPosition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"correct": correct})
}
