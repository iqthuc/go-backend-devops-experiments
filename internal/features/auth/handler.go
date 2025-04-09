package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	messages "github.com/iqthuc/sport-shop/internal"
	"github.com/iqthuc/sport-shop/internal/delivery/http/middleware"
	"github.com/iqthuc/sport-shop/utils"
)

type Handler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) Handler {
	return &handler{
		useCase: useCase,
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
	if err := h.useCase.RegisterUser(r.Context(), user); err != nil {
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
	result, status, err := h.useCase.Login(r.Context(), req)
	if err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, messages.LoginFailedInternal)
		return
	}
	if status == LoginStatusUserNotFound || status == LoginStatusInvalidPassword {
		utils.ErrorJsonResponse(w, http.StatusUnauthorized, messages.LoginIncorrectCredentials)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    result.RefreshToken,
		Path:     "/api/products",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int((24 * time.Hour).Seconds()),
	})

	utils.SuccessJsonResponse(w, result.AccessToken, messages.LoginSuccesfully)
}

func (h *handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.ContextKey.UserID)
	if userId == nil {
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "User not found")
		return
	}
	user, err := h.useCase.GetUserInfo(r.Context(), userId.(int64))
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

func (h *handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req refreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorJsonResponse(w, http.StatusBadRequest, messages.InvalidRequest)
		return
	}
	newAccessToken, err := h.useCase.RefreshToken(r.Context(), req.Token)
	if err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Failed")
		return
	}
	utils.SuccessJsonResponse(w, newAccessToken, messages.LoginSuccesfully)
}
