package domain

import "context"

type Stage struct {
	ID           string
	Name         string
	StageNumber  int
	QuestionType string
}

type StageRepository interface {
	FindAll(ctx context.Context) ([]*Stage, error)
	FindByID(ctx context.Context, id string) (*Stage, error)
}

type StageUsecase interface {
	GetStages(ctx context.Context) ([]*Stage, error)
	GetStageByID(ctx context.Context, id string) (*Stage, error)
}
