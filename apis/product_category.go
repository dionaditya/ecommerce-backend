package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/ecommerce-backend/models"
	httpUtil "github.com/masterraf21/ecommerce-backend/utils/http"
)

type productCategoryAPI struct {
	ProductCategoryUsecase models.ProductCategoryUsecase
}

// NewProductAPI will create api for product
func NewProductCategoryAPI(r *mux.Router, pru models.ProductCategoryUsecase) {
	productCategoryAPI := &productCategoryAPI{
		ProductCategoryUsecase: pru,
	}

	r.HandleFunc("/product_category", productCategoryAPI.Create).Methods("POST")
	r.HandleFunc("/product_category", productCategoryAPI.GetAll).Methods("GET")
	r.HandleFunc("/product_category/{id_product_category}", productCategoryAPI.GetByID).Methods("GET")
	r.HandleFunc("/product_category/{id_product_category}", productCategoryAPI.UpdateProductCategory).Methods("PUT")
	r.HandleFunc("/product_category/{id_product_category}", productCategoryAPI.DeleteCategory).Methods("DELETE")
}

func (p *productCategoryAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.ProductCategoryBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	id, err := p.ProductCategoryUsecase.CreateProductCategory(body)

	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to creat product category", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"category_id"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (p *productCategoryAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := p.ProductCategoryUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get product data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.ProductCategory `json:"data"`
	}
	data.Data = result
	httpUtil.HandleJSONResponse(w, r, data)
}

func (p *productCategoryAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productCategoryID, err := strconv.ParseInt(params["id_product_category"], 10, 64)

	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_product_category"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := p.ProductCategoryUsecase.GetByID(uint32(productCategoryID))

	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get product data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.ProductCategory `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)

}

func (p *productCategoryAPI) UpdateProductCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productCategoryID, err := strconv.ParseInt(params["id_product_category"], 10, 64)

	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_product_category"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	var body models.ProductCategoryBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	result, err := p.ProductCategoryUsecase.UpdateProductCategory(uint32(productCategoryID), body.CategoryName)

	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to update product category", http.StatusInternalServerError)
		return
	}

	httpUtil.HandleJSONResponse(w, r, result)

}

func (p *productCategoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productCategoryID, err := strconv.ParseInt(params["id_product_category"], 10, 64)

	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_product_category"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = p.ProductCategoryUsecase.DeleteCategory(uint32(productCategoryID))

	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to delete product category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
