package user

import (
	"net/http"

	"github.com/giomhern/ecomm/types"
	"github.com/giomhern/ecomm/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// accept JSON paywall --> check if user exists (if not we create user)

	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
