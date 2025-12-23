package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGetUserProtected(t *testing.T) {
	s := newTestServer()

	// 1. Setup: Create a user and get a token
	email := fmt.Sprintf("user_test_%d@example.com", time.Now().UnixNano())
	password := "password123"

	// Register
	regPayload, _ := json.Marshal(map[string]string{"email": email, "password": password})
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(regPayload))
	req.Header.Set("Content-Type", "application/json")
	executeRequest(req, s)

	// Login
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(regPayload))
	req.Header.Set("Content-Type", "application/json")
	resp := executeRequest(req, s)

	var loginResp map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &loginResp)
	if loginResp["token"] == nil {
		t.Fatalf("Failed to login, token is nil. Body: %s", resp.Body.String())
	}
	token := loginResp["token"].(string)

	// 2. Test: Get User with Token
	req, _ = http.NewRequest("GET", "/users?email="+email, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response)

	var userResp map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &userResp)

	if userResp["Email"] != email {
		t.Errorf("Expected email %s, got %v", email, userResp["Email"])
	}
}

func TestGetUserUnauthorized(t *testing.T) {
	s := newTestServer()

	req, _ := http.NewRequest("GET", "/users?email=some@email.com", nil)
	// No Authorization header

	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusUnauthorized, response)
}
