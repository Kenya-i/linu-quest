package repository

import (
	"context"

	"github.com/Kenya-i/linu-quest/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type questionRepository struct {
	db *pgxpool.Pool
}

func NewQuestionRepository(db *pgxpool.Pool) domain.QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) FindByID(ctx context.Context, id string) (*domain.Question, error) {
	query := `SELECT ID, STAGE_ID, QUESTION_NUMBER, QUESTION_TEXT, ANSWER_COMMAND, CORRECT_POSITION FROM QUESTIONS WHERE ID = $1`
	var question domain.Question
	err := r.db.QueryRow(ctx, query, id).Scan(
		&question.ID,
		&question.StageID,
		&question.QuestionNumber,
		&question.QuestionText,
		&question.AnswerCommand,
		&question.CorrectPosition)
	if err != nil {
		return nil, err
	}

	choicesQuery := `SELECT CHOICE_TEXT FROM QUESTION_CHOICES WHERE QUESTION_ID = $1 ORDER BY POSITION`
	rows, err := r.db.Query(ctx, choicesQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var choiceText string
		if err := rows.Scan(&choiceText); err != nil {
			return nil, err
		}
		question.Choices = append(question.Choices, choiceText)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &question, nil
}

func (r *questionRepository) FindByStageID(ctx context.Context, stageID string) ([]*domain.Question, error) {
	query := `SELECT ID, STAGE_ID, QUESTION_NUMBER, QUESTION_TEXT, ANSWER_COMMAND, CORRECT_POSITION FROM QUESTIONS WHERE STAGE_ID = $1 ORDER BY QUESTION_NUMBER`
	rows, err := r.db.Query(ctx, query, stageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*domain.Question
	for rows.Next() {
		var question domain.Question
		err := rows.Scan(
			&question.ID,
			&question.StageID,
			&question.QuestionNumber,
			&question.QuestionText,
			&question.AnswerCommand,
			&question.CorrectPosition)
		if err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// 1. ID一覧を抽出
	questionIDs := make([]string, len(questions))
	for i, q := range questions {
		questionIDs[i] = q.ID
	}

	// 2. まとめて選択肢を取得
	choicesQuery := `SELECT QUESTION_ID, CHOICE_TEXT FROM QUESTION_CHOICES WHERE QUESTION_ID = ANY($1) ORDER BY POSITION`
	choiceRows, err := r.db.Query(ctx, choicesQuery, questionIDs)
	if err != nil {
		return nil, err
	}
	defer choiceRows.Close()

	// 3. question_idごとに振り分け
	choicesMap := make(map[string][]string)
	for choiceRows.Next() {
		var questionID, choiceText string
		if err := choiceRows.Scan(&questionID, &choiceText); err != nil {
			return nil, err
		}
		choicesMap[questionID] = append(choicesMap[questionID], choiceText)
	}
	if err := choiceRows.Err(); err != nil {
		return nil, err
	}

	// 4. 各questionにChoicesを割り当て
	for _, q := range questions {
		q.Choices = choicesMap[q.ID]
	}

	return questions, nil
}
