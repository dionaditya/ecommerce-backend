package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductCategory struct {
	ID           uint32             `bson:"category_id" json:"category_id"`
	CategoryName string             `bson:"category_name" json:"category_name"`
	OID          primitive.ObjectID `bson:"_id" json:"OID"`
}

type ProductCategoryResponse struct {
	Data  []ProductCategory `json:"data"`
	Total int               `json:"total"`
}
type ProductCategoryBody struct {
	CategoryName string `bson:"category_name" json:"category_name"`
}

type ProductCategoryRepository interface {
	Store(productCategory *ProductCategory) (uint32, error)
	GetAll(searchQuery string, page string, limit string) (*ProductCategoryResponse, error)
	GetByID(id uint32) (*ProductCategory, error)
	UpdateArbitrary(id uint32, value interface{}) (*ProductCategory, error)
	DeleteCategory(id uint32) error
}

type ProductCategoryUsecase interface {
	CreateProductCategory(productCategory ProductCategoryBody) (uint32, error)
	GetAll(searchQuery string, page string, limit string) (*ProductCategoryResponse, error)
	GetByID(id uint32) (*ProductCategory, error)
	UpdateProductCategory(id uint32, value string) (*ProductCategory, error)
	DeleteCategory(id uint32) error
}
