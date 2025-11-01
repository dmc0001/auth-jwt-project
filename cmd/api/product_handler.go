package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dmc0001/auth-jwt-project/internal/types"
	"github.com/dmc0001/auth-jwt-project/internal/utils"
)

func (app *Application) getProduct(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("param")

	id, err := strconv.Atoi(param)
	if err != nil {
		// Not a number → treat as name
		app.getProductByName(w, param)
		return
	}

	app.getProductById(w, id)

}

func (app *Application) getProducts(w http.ResponseWriter, r *http.Request) {

	products, err := app.config.productModel.GetProducts()
	if err != nil {
		utils.WritingError(w, http.StatusNotFound, err)
		return
	}

	err = utils.ParsingToJson(w, http.StatusOK, products)
	if err != nil {
		utils.WritingError(w, http.StatusInternalServerError, err)
	}

}

func (app *Application) getProductById(w http.ResponseWriter, id int) {

	product, err := app.config.productModel.GetProductById(id)
	if err != nil {
		utils.WritingError(w, http.StatusNotFound, err)
		return
	}

	err = utils.ParsingToJson(w, http.StatusOK, product)
	if err != nil {
		utils.WritingError(w, http.StatusInternalServerError, err)
	}

}
func (app *Application) getProductByName(w http.ResponseWriter, name string) {

	products, err := app.config.productModel.GetProductByName(name)
	if err != nil {
		utils.WritingError(w, http.StatusNotFound, err)
		return
	}

	err = utils.ParsingToJson(w, http.StatusOK, products)
	if err != nil {
		utils.WritingError(w, http.StatusInternalServerError, err)
	}

}

func (app *Application) createProduct(w http.ResponseWriter, r *http.Request) {
	var product types.CreateProductRequest

	if err := utils.ParsingFromJson(r, &product); err != nil {
		if strings.Contains(err.Error(), "Cannot unmarshal") {
			utils.WritingError(w, http.StatusBadRequest,
				fmt.Errorf("Invalid field types in request body. Ensure numeric fields like 'price' and 'quantity' are numbers, not strings."))
			return
		}
		utils.WritingError(w, http.StatusBadRequest, err)
		return
	}

	err := app.config.productModel.CreateProduct(product)
	if err != nil {
		utils.WritingError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.ParsingToJson(w, http.StatusOK, map[string]string{"Message": "✅ Product created successfully"})
	if err != nil {
		utils.WritingError(w, http.StatusInternalServerError, err)
	}

}
