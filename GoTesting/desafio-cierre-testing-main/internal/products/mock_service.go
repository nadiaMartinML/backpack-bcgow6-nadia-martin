package products

import "fmt"

type MockService struct {
	DataMock []Product
	Error    string
}

func (m *MockService) GetAllBySeller(sellerID string) ([]Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}

	return m.DataMock, nil
}
