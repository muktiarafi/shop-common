package common

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareRequireRoles(t *testing.T) {
	t.Run("test admin role", func(t *testing.T) {
		userPayload := &UserPayload{
			ID:    1,
			Email: "bambank@gmail.com",
			Roles: []string{"USER", "ADMIN"},
		}
		token, _ := CreateToken(userPayload)
		req := httptest.NewRequest(http.MethodGet, "/admins", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()

		mux.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expecting status to be %d, but got %d instead", http.StatusOK, res.Code)
		}
	})

	t.Run("test admin route without valid role", func(t *testing.T) {
		userPayload := &UserPayload{
			ID:    1,
			Email: "bambank@gmail.com",
			Roles: []string{"USER"},
		}
		token, _ := CreateToken(userPayload)
		req := httptest.NewRequest(http.MethodGet, "/admins", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()

		mux.ServeHTTP(res, req)

		if res.Code != http.StatusUnauthorized {
			t.Errorf("expecting status to be %d, but got %d instead", http.StatusUnauthorized, res.Code)
		}
	})

	t.Run("test admin or seller role", func(t *testing.T) {
		userPayload := &UserPayload{
			ID:    1,
			Email: "bambank@gmail.com",
			Roles: []string{"USER", "SELLER"},
		}
		token, _ := CreateToken(userPayload)
		req := httptest.NewRequest(http.MethodGet, "/sellers", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()

		mux.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expecting status to be %d, but got %d instead", http.StatusOK, res.Code)
		}
	})
}
