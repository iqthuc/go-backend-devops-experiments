package product

import (
	"fmt"
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
	products, err := h.userCase.GetProducts(r.Context())

	if err != nil {
		utils.JsonResponse(w, utils.APIResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Status:  0,
		})
	}

	utils.JsonResponse(w, utils.APIResponse{
		Status:  1,
		Message: "ok",
		Data:    products,
		Code:    http.StatusOK,
	})

}
func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create product")
}
