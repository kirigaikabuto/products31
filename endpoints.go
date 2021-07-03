package products31

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ProductHttpEndpoints interface {
	CreateProduct() func(w http.ResponseWriter, r *http.Request)
	ListProduct() func(w http.ResponseWriter, r *http.Request)
}

type productHttpEndpoints struct {
	productStore ProductStore
}

func NewProductHttpEndpoints(p ProductStore) ProductHttpEndpoints {
	return &productHttpEndpoints{productStore: p}
}

func (p *productHttpEndpoints) CreateProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		cmd := &CreateProductCommand{}
		err = json.Unmarshal(jsonData, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		response, err := p.productStore.Create(&Product{
			Name:        cmd.Name,
			Description: cmd.Description,
			Price:       cmd.Price,
		})
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusCreated, response)
		return
	}
}

func (p *productHttpEndpoints) ListProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := p.productStore.List()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusCreated, products)
		return
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
