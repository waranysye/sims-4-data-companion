package repository

type MockCareerRepository struct {
	MockData []CareerRecommendation
	MockErr  error
}

func (m *MockCareerRepository) GetRecommendations() ([]CareerRecommendation, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockData, nil
}

func (m *MockCareerRepository) CreateRecommendation(careerID, traitID, compatibilityScore int, reason string) error {
	return m.MockErr
}

func (m *MockCareerRepository) UpdateRecommendation(careerID, traitID, compatibilityScore int, reason string) error {
	return m.MockErr
}

func (m *MockCareerRepository) DeleteRecommendation(careerID, traitID int) error {
	return m.MockErr
}