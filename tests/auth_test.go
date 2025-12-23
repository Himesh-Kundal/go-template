package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestRegisterAndLogin(t *testing.T) {
	s := newTestServer()

	// Unique email for this test run
	email := fmt.Sprintf("auth_test_%d@example.com", time.Now().UnixNano())
	password := "securepass"

	// 1. Register
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusCreated, response)

	var userResp map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &userResp)

	if userResp["Email"] != email {
		t.Errorf("Expected email %s, got %v", email, userResp["Email"])
	}

	// 2. Login
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response)

	var loginResp map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &loginResp)

	if loginResp["token"] == "" {
		t.Error("Expected JWT token in login response")
	}
}

func TestLoginInvalidCredentials(t *testing.T) {
	s := newTestServer()

	payload := map[string]string{
		"email":    "nonexistent@example.com",
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusUnauthorized, response)
}
