package http

import (
	"clean-architecture/internal/domain"
	"clean-architecture/pkg/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(usecase domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		response.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.usecase.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
		} else {
			response.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response.Success(w, "user retrieved successfully", user)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "users retrieved successfully", users)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.usecase.Create(&user); err != nil {
		if errors.Is(err, domain.ErrNameRequired) || errors.Is(err, domain.ErrEmailRequired) {
			response.Error(w, http.StatusBadRequest, err.Error())
		} else {
			response.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response.Created(w, "user created successfully", user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		response.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.usecase.Update(id, &user); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
		} else if errors.Is(err, domain.ErrNameRequired) || errors.Is(err, domain.ErrEmailRequired) {
			response.Error(w, http.StatusBadRequest, err.Error())
		} else {
			response.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response.Success(w, "user updated successfully", user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		response.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
		} else {
			response.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response.Success(w, "user deleted successfully", nil)
}
