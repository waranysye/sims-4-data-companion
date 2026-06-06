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
}

type CareerRepository struct {
	DB *pgxpool.Pool
}

func NewCareerRepository(db *pgxpool.Pool) *CareerRepository {
	return &CareerRepository{DB: db}
}

func (r *CareerRepository) GetRecommendations() ([]CareerRecommendation, error) {
	// 👈 Query SQL kita perbarui untuk menarik semua kolom dari careers dan career_recommendations
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
		// 👈 Scan semua data baru ke dalam struct variabel masing-masing
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
