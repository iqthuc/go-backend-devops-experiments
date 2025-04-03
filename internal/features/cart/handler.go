package cart

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/iqthuc/sport-shop/utils"
)

type Handler struct {
	userCase UserCase
}

func NewHandler(userCase UserCase) *Handler {
	return &Handler{userCase: userCase}
}

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := utils.ParseQueryParam[int64](r, "user_id")

	cart, err := h.userCase.GetCart(r.Context(), userID)
	if err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Failed to get cart")
		return
	}

	if cart == nil {
		utils.SuccessJsonResponse(w, nil, "Cart is empty")
		return
	}

	utils.SuccessJsonResponse(w, cart, "Cart retrieved successfully")
}

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req AddToCartRequestParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}
	item := &CartItem{
		ProductID: req.ProductID,
		Name:      req.Name,
		Price:     req.Price,
		Quantity:  req.Quantity,
	}
	if err := h.userCase.AddToCart(r.Context(), req.UserID, item); err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	utils.SuccessJsonResponse(w, nil, "Add product to cart successfully")
}
