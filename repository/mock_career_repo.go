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