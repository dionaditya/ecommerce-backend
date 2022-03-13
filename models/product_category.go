package models

type ProductCategory struct {
	ID           uint32 `bson:"category_id" json:"category_id"`
	CategoryName string `bson:"category_name" json:"category_name"`
	OID          string `bson:"_id" json:"OID"`
}

type ProductCategoryBody struct {
	CategoryName string `bson:"category_name" json:"category_name"`
}

type ProductCategoryRepository interface {
	Store(productCategory *ProductCategory) (uint32, error)
	GetAll() ([]ProductCategory, error)
	GetByID(id uint32) (*ProductCategory, error)
	UpdateArbitrary(id uint32, value interface{}) (*ProductCategory, error)
	DeleteCategory(id uint32) error
}

type ProductCategoryUsecase interface {
	CreateProductCategory(productCategory ProductCategoryBody) (uint32, error)
	GetAll() ([]ProductCategory, error)
	GetByID(id uint32) (*ProductCategory, error)
	UpdateProductCategory(id uint32, value string) (*ProductCategory, error)
	DeleteCategory(id uint32) error
}
