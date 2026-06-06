package usecase

import (
	"testing"
)

func TestContohValidasiSkorKecocokan(t *testing.T) {
	// Ekspektasi skor maksimal di platform data kita adalah 5
	maxScore := 5
	currentScore := 5

	if currentScore > maxScore {
		t.Errorf("Error: Skor kecocokan tidak boleh lebih dari %d, mendapatkan %d", maxScore, currentScore)
	}
}