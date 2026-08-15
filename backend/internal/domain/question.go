package domain

import "context"

type Question struct {
	ID              string
	StageID         string
	QuestionNumber  int
	QuestionText    string
	AnswerCommand   string
	Choices         []string
	CorrectPosition int
}

type QuestionRepository interface {
	FindByStageID(ctx context.Context, stageID string) ([]*Question, error)
	FindByID(ctx context.Context, id string) (*Question, error)
}

type QuestionUsecase interface {
	GetQuestionsByStageID(ctx context.Context, stageID string) ([]*Question, error)
	CheckAnswer(ctx context.Context, questionID string, selectedPosition int) (bool, error)
}
