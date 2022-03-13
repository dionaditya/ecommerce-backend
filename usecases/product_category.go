package usecases

import (
	"github.com/masterraf21/ecommerce-backend/models"
)

type productCategoryUsecase struct {
	Repo models.ProductCategoryRepository
}

// NewProductUsecase will initiate usecase
func NewProductCategoryUseCase(productCategoryRepo models.ProductCategoryRepository) models.ProductCategoryUsecase {
	return &productCategoryUsecase{Repo: productCategoryRepo}
}

func (u *productCategoryUsecase) CreateProductCategory(body models.ProductCategoryBody) (id uint32, err error) {
	productCategory := models.ProductCategory{
		CategoryName: body.CategoryName,
	}

	_id, err := u.Repo.Store(&productCategory)

	if err != nil {
		return
	}

	id = _id

	return
}

func (u *productCategoryUsecase) GetAll(searchQuery string, page string, limit string) (res *models.ProductCategoryResponse, err error) {
	res, err = u.Repo.GetAll(searchQuery, page, limit)
	return
}

func (u *productCategoryUsecase) GetByID(id uint32) (res *models.ProductCategory, err error) {
	res, err = u.Repo.GetByID(id)

	return
}

func (u *productCategoryUsecase) UpdateProductCategory(id uint32, value string) (res *models.ProductCategory, err error) {
	res, err = u.Repo.UpdateArbitrary(id, value)

	return
}

func (u *productCategoryUsecase) DeleteCategory(id uint32) error {
	err := u.Repo.DeleteCategory(id)

	return err
}
