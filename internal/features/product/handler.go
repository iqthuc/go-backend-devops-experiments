package product

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iqthuc/sport-shop/utils"
)

type handler struct {
	userCase UserCase
}

func NewHandler(service UserCase) *handler {
	return &handler{
		userCase: service,
	}
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	requestParams := GetProductsRequestParams{
		page:       utils.ParseQueryInt(r, "page", 1),
		searchKey:  utils.ParseQueryString(r, "search_key"),
		categoryId: utils.ParseQueryInt(r, "category_id", 0),
	}

	result, err := h.userCase.GetProducts(r.Context(), requestParams)

	if err != nil {
		log.Printf("error: %v", err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Oops! Something went wrong")
		return
	}

	utils.SuccessJsonResponse(w, result, "ok")
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create product")
}
