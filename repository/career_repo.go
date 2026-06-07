package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// 👈 Kita tambahkan field baru di struct agar Golang bisa menampung data barunya
type CareerRecommendation struct {
	CareerName         string `json:"career_name"`
	Branch             string `json:"branch"`
	BaseSalary         int    `json:"base_salary"`
	IdealMood          string `json:"ideal_mood"`
	TraitName          string `json:"trait_name"`
	CompatibilityScore int    `json:"compatibility_score"`
	Reason             string `json:"reason"`
}

type CareerRepositoryInterface interface {
	GetRecommendations() ([]CareerRecommendation, error)
	CreateRecommendation(careerID, traitID, compatibilityScore int, reason string) error
	UpdateRecommendation(careerID, traitID, compatibilityScore int, reason string) error
	DeleteRecommendation(careerID, traitID int) error
}

type CareerRepository struct {
	DB *pgxpool.Pool
}

func NewCareerRepository(db *pgxpool.Pool) *CareerRepository {
	return &CareerRepository{DB: db}
}

func (r *CareerRepository) GetRecommendations() ([]CareerRecommendation, error) {
	query := `
		SELECT 
			c.name, c.branch, c.base_salary, c.ideal_mood, 
			t.name, cr.compatibility_score, cr.reason
		FROM career_recommendations cr
		JOIN careers c ON cr.career_id = c.id
		JOIN traits t ON cr.trait_id = t.id
		ORDER BY cr.compatibility_score DESC;
	`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []CareerRecommendation
	for rows.Next() {
		var row CareerRecommendation
		err := rows.Scan(
			&row.CareerName, &row.Branch, &row.BaseSalary, &row.IdealMood,
			&row.TraitName, &row.CompatibilityScore, &row.Reason,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, row)
	}
	return list, nil
}

func (r *CareerRepository) CreateRecommendation(careerID, traitID, compatibilityScore int, reason string) error {
	query := `INSERT INTO career_recommendations (career_id, trait_id, compatibility_score, reason) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(context.Background(), query, careerID, traitID, compatibilityScore, reason)
	return err
}

func (r *CareerRepository) UpdateRecommendation(careerID, traitID, compatibilityScore int, reason string) error {
	query := `UPDATE career_recommendations SET compatibility_score = $1, reason = $2 WHERE career_id = $3 AND trait_id = $4`
	_, err := r.DB.Exec(context.Background(), query, compatibilityScore, reason, careerID, traitID)
	return err
}

func (r *CareerRepository) DeleteRecommendation(careerID, traitID int) error {
	query := `DELETE FROM career_recommendations WHERE career_id = $1 AND trait_id = $2`
	_, err := r.DB.Exec(context.Background(), query, careerID, traitID)
	return err
}
