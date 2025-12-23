package tests

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go-template/internal/auth"
	"go-template/internal/db"
	"go-template/internal/server"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var testQueries *db.Queries
var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	// Only load .env if DATABASE_URL is NOT set.
	// This allows 'make test' to override it with localhost.
	if os.Getenv("DATABASE_URL") == "" {
		_ = godotenv.Load("../.env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	log.Printf("Test DATABASE_URL: %s", dbURL)

	if dbURL == "" {
		log.Println("Skipping integration tests: DATABASE_URL not set")
		os.Exit(0)
	}

	ctx := context.Background()
	var err error
	testPool, err = pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer testPool.Close()

	testQueries = db.New(testPool)

	// Set a secret for testing auth
	auth.SetSecret("test-secret")

	code := m.Run()

	os.Exit(code)
}

func newTestServer() *server.Server {
	return server.New(testQueries)
}

func executeRequest(req *http.Request, s *server.Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected int, rr *httptest.ResponseRecorder) {
	if expected != rr.Code {
		t.Errorf("Expected response code %d. Got %d. Body: %s\n", expected, rr.Code, rr.Body.String())
	}
}
