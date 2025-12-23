package handlers

import (
	"encoding/json"
	"net/http"

	"go-template/internal/auth"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := h.q.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if !auth.CheckPasswordHash(req.Password, user.PasswordHash) {
		respondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := auth.GenerateToken(uuidToString(user.ID), user.Email)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not generate token")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{
		"token":   token,
		"user_id": uuidToString(user.ID),
		"email":   user.Email,
	})
}
