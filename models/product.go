package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product represents product
type Product struct {
	ID                uint32  `bson:"id_product" json:"id_product"`
	ProductName       string  `bson:"product_name" json:"product_name"`
	Description       string  `bson:"description" json:"description"`
	Price             float32 `bson:"price" json:"price"`
	SellerID          uint32  `bson:"id_seller" json:"id_seller"`
	Seller            *Seller `bson:"seller" json:"seller"`
	ProductCategoryID uint32  `bson:"product_category_id" json:"product_category_id"`
}

// ProductBody for receiving body grom json
type ProductBody struct {
	ProductName       string  `json:"product_name"`
	Description       string  `json:"description"`
	Price             float32 `json:"price"`
	SellerID          uint32  `json:"id_seller"`
	ProductCategoryID uint32  `json:"product_category_id"`
}
type ProductResponse struct {
	Data  []Product
	Total int
}

// ProductRepository represents repo functions for product
type ProductRepository interface {
	Store(product *Product) (primitive.ObjectID, error)
	GetAll(searchQuery string, page string, limit string, category_id string, sortBy string, sortDirection string) (*ProductResponse, error)
	GetBySellerID(sellerID uint32) ([]Product, error)
	GetByID(id uint32) (*Product, error)
	GetByOID(oid primitive.ObjectID) (*Product, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
}

// ProductUsecase usecase for product
type ProductUsecase interface {
	CreateProduct(product ProductBody) (uint32, error)
	GetAll(searchQuery string, page string, limit string, category_id string, sortBy string, sortDirection string) (*ProductResponse, error)
	GetBySellerID(sellerID uint32) ([]Product, error)
	GetByID(id uint32) (*Product, error)
}
