package repository

import (
	"context"

	"github.com/Kenya-i/linu-quest/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type stageRepository struct {
	db *pgxpool.Pool
}

func NewStageRepository(db *pgxpool.Pool) domain.StageRepository {
	return &stageRepository{db: db}
}

func (r *stageRepository) FindAll(ctx context.Context) ([]*domain.Stage, error) {
	query := `SELECT ID, NAME, STAGE_NUMBER, QUESTION_TYPE FROM STAGES`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []*domain.Stage
	for rows.Next() {
		var stage domain.Stage
		err := rows.Scan(
			&stage.ID,
			&stage.Name,
			&stage.StageNumber,
			&stage.QuestionType)
		if err != nil {
			return nil, err
		}
		stages = append(stages, &stage)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stages, nil
}

func (r *stageRepository) FindByID(ctx context.Context, id string) (*domain.Stage, error) {
	query := `SELECT ID, NAME, STAGE_NUMBER, QUESTION_TYPE FROM STAGES WHERE ID = $1`
	var stage domain.Stage
	err := r.db.QueryRow(ctx, query, id).Scan(
		&stage.ID,
		&stage.Name,
		&stage.StageNumber,
		&stage.QuestionType)

	if err != nil {
		return nil, err
	}

	return &stage, nil
}
