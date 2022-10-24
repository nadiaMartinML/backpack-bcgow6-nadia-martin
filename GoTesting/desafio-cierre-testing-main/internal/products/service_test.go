package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// arrange
	product := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	myMock := MockService{
		DataMock: product,
		Error:    "",
	}

	service := NewService(&myMock)
	sellerToSearch := "FEX112AC"

	// act
	result, err := service.GetAllBySeller(sellerToSearch)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestGetAllSellerIdIncorrect(t *testing.T) {
	// arrage
	errExpect := "error in repository"
	myMock := MockService{
		DataMock: nil,
		Error:    "error in repository",
	}

	service := NewService(&myMock)

	// act
	result, err := service.GetAllBySeller("")

	// assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errExpect)
	assert.Nil(t, result)
}
