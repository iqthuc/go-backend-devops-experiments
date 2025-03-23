package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/iqthuc/sport-shop/utils"
)

type Handler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	userCase UserCase
}

func NewHandler(userCase UserCase) Handler {
	return &handler{
		userCase: userCase,
	}
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user := &User{
		username:     req.Username,
		email:        req.Email,
		passwordHash: req.Password,
		fullName:     req.FullName,
		phoneNumber:  req.PhoneNumber,
	}
	if err := h.userCase.RegisterUser(r.Context(), user); err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Something errors")
		return
	}
	utils.SuccessJsonResponse(w, nil, "User registered successfully")

}
