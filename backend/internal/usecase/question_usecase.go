package usecase

import (
	"context"

	"github.com/Kenya-i/linu-quest/internal/domain"
)

type questionUsecase struct {
	questionRepo domain.QuestionRepository
}

func NewQuestionUsecase(questionRepo domain.QuestionRepository) domain.QuestionUsecase {
	return &questionUsecase{questionRepo: questionRepo}
}

func (u *questionUsecase) GetQuestionsByStageID(ctx context.Context, stageID string) ([]*domain.Question, error) {
	return u.questionRepo.FindByStageID(ctx, stageID)
}

func (u *questionUsecase) CheckAnswer(ctx context.Context, questionID string, selectedPosition int) (bool, error) {
	question, err := u.questionRepo.FindByID(ctx, questionID)
	if err != nil {
		return false, err
	}

	return question.CorrectPosition == selectedPosition, nil
}
