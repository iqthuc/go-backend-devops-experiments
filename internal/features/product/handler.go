package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/iqthuc/sport-shop/utils"
)

type Handler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductDetails(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) Handler {
	return &handler{
		useCase: useCase,
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

	result, err := h.useCase.GetProducts(r.Context(), requestParams)

	if err != nil {
		log.Printf("error: %v", err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Oops! Something went wrong")
		return
	}

	utils.SuccessJsonResponse(w, result, "ok")
}
func (h *handler) GetProductDetails(w http.ResponseWriter, r *http.Request) {
	product_id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Something errors")
		return
	}
	productDetails, err := h.useCase.GetProductDetails(r.Context(), product_id)
	if err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Failed to fetch products")
		return
	}
	utils.SuccessJsonResponse(w, productDetails, "ok")
}
