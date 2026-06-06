package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Struktur Data untuk menampung hasil query gabungan (JOIN)
type CareerRecommendation struct {
	CareerName         string `json:"career_name"`
	Branch             string `json:"branch"`
	BaseSalary         int    `json:"base_salary"`
	IdealMood          string `json:"ideal_mood"`
	TraitName          string `json:"trait_name"`
	CompatibilityScore int    `json:"compatibility_score"`
	Reason             string `json:"reason"`
}

type CareerRepository struct {
	DB *pgxpool.Pool
}

func NewCareerRepository(db *pgxpool.Pool) *CareerRepository {
	return &CareerRepository{DB: db}
}

// Fungsi untuk mengambil rekomendasi kecocokan karir dan sifat
func (r *CareerRepository) GetRecommendations() ([]CareerRecommendation, error) {
	query := `
		SELECT 
			c.name, c.branch, c.base_salary, c.ideal_mood,
			t.name, ctr.compatibility_score, ctr.reason
		FROM career_trait_recommendations ctr
		JOIN careers c ON ctr.career_id = c.id
		JOIN traits t ON ctr.trait_id = t.id
		ORDER BY ctr.compatibility_score DESC;
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []CareerRecommendation
	for rows.Next() {
		var cr CareerRecommendation
		err := rows.Scan(
			&cr.CareerName, &cr.Branch, &cr.BaseSalary, &cr.IdealMood,
			&cr.TraitName, &cr.CompatibilityScore, &cr.Reason,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, cr) //  Benar
	}

	return list, nil
}
