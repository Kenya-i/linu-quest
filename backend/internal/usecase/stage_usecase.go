package usecase

import (
	"context"

	"github.com/Kenya-i/linu-quest/internal/domain"
)

type stageUsecase struct {
	stageRepo domain.StageRepository
}

func NewStageUsecase(stageRepo domain.StageRepository) domain.StageUsecase {
	return &stageUsecase{stageRepo: stageRepo}
}

func (u *stageUsecase) GetStages(ctx context.Context) ([]*domain.Stage, error) {
	return u.stageRepo.FindAll(ctx)
}

func (u *stageUsecase) GetStageByID(ctx context.Context, id string) (*domain.Stage, error) {
	return u.stageRepo.FindByID(ctx, id)
}
