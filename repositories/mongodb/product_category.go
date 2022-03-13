package mongodb

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	_ "fmt"
)

type productCategoryRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewOrderRepo will initiate order repository object
func NewProductCategoryRepo(instance *mongo.Database, ctr models.CounterRepository) models.ProductCategoryRepository {
	return &productCategoryRepo{Instance: instance, CounterRepo: ctr}
}

func (r *productCategoryRepo) Store(productCategory *models.ProductCategory) (id uint32, err error) {
	collectionName := "product_category"
	identifier := "id_product_category"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	category_id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	productCategory.ID = category_id

	result, err := collection.InsertOne(ctx, productCategory)

	_ = result

	if err != nil {
		return
	}

	id = category_id

	return
}

func (r *productCategoryRepo) GetByID(id uint32) (res *models.ProductCategory, err error) {
	collection := r.Instance.Collection("product_category")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"category_id": id}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *productCategoryRepo) GetAll() (res []models.ProductCategory, error error) {
	collection := r.Instance.Collection("product_category")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.ProductCategory, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *productCategoryRepo) UpdateArbitrary(id uint32, value interface{}) (*models.ProductCategory, error) {
	collection := r.Instance.Collection("product_category")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"category_id": id},
		bson.M{"$set": bson.M{"category_name": value}},
	)

	var productCategory *models.ProductCategory

	err = collection.FindOne(ctx, bson.M{"category_id": id}).Decode(&productCategory)

	if err != nil {
		return productCategory, err
	}

	return productCategory, nil
}

func (r *productCategoryRepo) DeleteCategory(id uint32) error {
	collection := r.Instance.Collection("product_category")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"category_id": id})

	fmt.Println(result.DeletedCount)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("No documents deleted")
	}

	return nil
}
