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
		page: utils.ParseQueryParam[int](r, "page"),
		Filters: Filters{
			Keyword:    utils.ParseQueryParam[string](r, "key_word"),
			CategoryID: utils.ParseQueryParam[int](r, "category_id"),
			PriceMin:   utils.ParseQueryParam[float64](r, "price_min"),
			PriceMax:   utils.ParseQueryParam[float64](r, "price_max"),
		},
		SortBy: SortBy{
			Field: utils.ParseQueryParam[string](r, "field"),
			Order: utils.ParseQueryParam[string](r, "order"),
		},
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
