package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	messages "github.com/iqthuc/sport-shop/internal"
	"github.com/iqthuc/sport-shop/internal/delivery/middleware"
	"github.com/iqthuc/sport-shop/utils"
)

type Handler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
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
		utils.ErrorJsonResponse(w, http.StatusBadRequest, messages.InvalidRequest)
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
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, messages.SomethingErrors)
		return
	}
	utils.SuccessJsonResponse(w, nil, messages.RegisterSuccessfully)

}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusBadRequest, messages.InvalidRequest)
		return
	}
	response, status, err := h.userCase.Login(r.Context(), req)
	if err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, messages.LoginFailedInternal)
		return
	}
	if status == LoginStatusUserNotFound || status == LoginStatusInvalidPassword {
		utils.ErrorJsonResponse(w, http.StatusUnauthorized, messages.LoginIncorrectCredentials)
		return
	}
	utils.SuccessJsonResponse(w, response, messages.LoginSuccesfully)
}

func (h *handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.ContextKey.UserID)
	if userId == nil {
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "User not found")
		return
	}
	user, err := h.userCase.GetUserInfo(r.Context(), userId.(int64))
	if err != nil {
		if errors.Is(err, Err.ErrUserNotFound) {
			utils.ErrorJsonResponse(w, http.StatusInternalServerError, "User not found")
			return
		}
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Failed to retrieve user profile")
		return
	}
	utils.SuccessJsonResponse(w, user, "User profile retrieved successfully")
}
