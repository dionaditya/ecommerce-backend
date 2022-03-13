package mongodb

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/models"
	"github.com/masterraf21/ecommerce-backend/utils/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type productRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewProductRepo will initiate product repo
func NewProductRepo(instance *mongo.Database, ctr models.CounterRepository) models.ProductRepository {
	return &productRepo{Instance: instance, CounterRepo: ctr}
}

func (r *productRepo) Store(product *models.Product) (oid primitive.ObjectID, err error) {
	collectionName := "product"
	identifier := "id_product"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	product.ID = id

	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return
	}

	_id := result.InsertedID
	oid = _id.(primitive.ObjectID)

	return
}

func (r *productRepo) GetAll(searchQuery string, page string, limit string, category_id string, sortBy string, sortDirection string) (res *models.ProductResponse, err error) {
	collection := r.Instance.Collection("product")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	query := bson.M{}

	query["product_name"] = primitive.Regex{
		Pattern: searchQuery,
		Options: "gi",
	}

	if category_id != "" {
		category_ID, _ := strconv.ParseInt(category_id, 10, 32)

		query["product_category_id"] = category_ID
	}

	var tempTotal []models.Product

	totalCursor, err := collection.Find(ctx, query)

	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = &models.ProductResponse{
				Data:  make([]models.Product, 0),
				Total: 0,
			}
			err = nil
			return
		}

		return
	}

	if err = totalCursor.All(ctx, &tempTotal); err != nil {
		return
	}

	var data []models.Product

	options := mongodb.GetOptions(page, limit, sortBy, sortDirection)

	cursor, err := collection.Find(ctx, query, options)

	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = &models.ProductResponse{
				Data:  make([]models.Product, 0),
				Total: 0,
			}
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &data); err != nil {
		return
	}

	res = &models.ProductResponse{
		Data:  data,
		Total: len(tempTotal),
	}

	return
}

func (r *productRepo) GetBySellerID(sellerID uint32) (res []models.Product, err error) {
	collection := r.Instance.Collection("product")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"id_seller": sellerID})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Product, 0)
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

func (r *productRepo) GetByID(id uint32) (res *models.Product, err error) {
	collection := r.Instance.Collection("product")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"id_product": id}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *productRepo) GetByOID(oid primitive.ObjectID) (res *models.Product, err error) {
	collection := r.Instance.Collection("product")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *productRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("product")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id_product": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}
