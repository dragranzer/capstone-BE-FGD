package bussiness_test

import (
	"os"
	"testing"

	"github.com/dragranzer/capstone-BE-FGD/features/categories"
	b_categories "github.com/dragranzer/capstone-BE-FGD/features/categories/bussiness"
	b_categories_mock "github.com/dragranzer/capstone-BE-FGD/features/categories/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	categoryData    b_categories_mock.Data
	categoryUsecase categories.Bussiness

	category []categories.Core
)

func TestMain(m *testing.M) {
	categoryUsecase = b_categories.NewCategoryBussiness(&categoryData)

	category = []categories.Core{
		{
			ID:   1,
			Name: "Hiburan",
		},
	}

	os.Exit(m.Run())
}

func TestAddCategory(t *testing.T) {
	t.Run("valid - add category", func(t *testing.T) {
		categoryData.On("InsertCategory", category[0]).Return(nil).Once()
		err := categoryUsecase.AddCategory(category[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - add category", func(t *testing.T) {
		categoryData.On("UpdateCategory", category[0]).Return(nil).Once()
		err := categoryUsecase.EditCategory(category[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete by id", func(t *testing.T) {
		categoryData.On("DeleteCategorybyId", category[0]).Return(nil).Once()
		err := categoryUsecase.DeleteCategorybyId(category[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get all category", func(t *testing.T) {
		categoryData.On("SelectAllCategory", category[0]).Return(category, nil)
		resp, err := categoryUsecase.GetAllCategory(category[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get category by id", func(t *testing.T) {
		categoryData.On("SelectCategorybyId", category[0]).Return(category[0], nil)
		resp, err := categoryUsecase.GetCategorybyId(category[0])

		assert.Equal(t, resp, category[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get category by name", func(t *testing.T) {
		categoryData.On("SelectCategorybyName", category[0]).Return(category[0], nil)
		resp, err := categoryUsecase.GetCategorybyName(category[0])

		assert.Equal(t, resp, category[0])
		assert.Equal(t, err, nil)
	})
}
