package common

import (
	"net/http"
	"os"
	"testing"
)

var mux *http.ServeMux

func TestMain(m *testing.M) {
	mux = http.NewServeMux()

	requireAdmin := RequireOneOfRole("ADMIN")
	requireAdminOrSeller := RequireOneOfRole("ADMIN", "SELLER")

	mux.Handle("/admins", RequireAuth(requireAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewResponse(http.StatusOK, "OK", "Congratulation you passed the middlewares!").SendJSON(w)
	}))))

	mux.Handle("/sellers", RequireAuth(requireAdminOrSeller(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewResponse(http.StatusOK, "OK", "Congratulation you passed the middlewares!").SendJSON(w)
	}))))

	os.Exit(m.Run())
}
