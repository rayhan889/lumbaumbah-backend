package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/types"
)

type mockUserStore struct{}

func TestAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	h := NewHanlder(&mockUserStore{})
	h.RegisterRoutes(r.Group("/"))

	r.Handle(http.MethodPost, "/signin", h.handleSignin)

	t.Run("payload can't be emtpy", func(t *testing.T) {
		payload := types.SigninPayload{}
		data, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("validation payload fails", func(t *testing.T) {
		payload := map[string]string{
			"email":      "test@test.com",
			"passworddd": "test",
		}
		data, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		payload := types.SigninPayload{
			Email:    "test@test.com",
			Password: "test",
		}
		data, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByEmail(email string) (types.User, error) {
	return types.User{}, nil
}

func (m *mockUserStore) GetUserByID(id string) (types.User, error) {
	return types.User{}, nil
}

func (m *mockUserStore) GetUserStatistics(uId string) (types.UserStatistics, error) {
	return types.UserStatistics{}, nil
}
