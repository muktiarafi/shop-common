package common

import (
	"errors"
	"net/http"
	"strings"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			exc := NewSingleMessageException(
				EUNAUTHORIZED,
				"RequiredAuth",
				"Not Authorized",
				errors.New("missing value in authorization header"),
			)
			SendError(w, exc)
			return
		}

		token, payload, err := ParseToken(authHeader[7:])
		if err != nil {
			exc := NewSingleMessageException(
				EUNAUTHORIZED,
				"RequiredAuth",
				"Not Authorized",
				errors.New("invalid token"),
			)
			SendError(w, exc)
			return
		}

		if !token.Valid {
			exc := NewSingleMessageException(
				EUNAUTHORIZED,
				"RequiredAuth",
				"Not Authorized",
				errors.New("invalid token"),
			)
			SendError(w, exc)
			return
		}

		ctx := UserPayloadToContext(payload, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireOneOfRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userPayload, err := UserPayloadFromContext("RequireRoles", r)
			if err != nil {
				SendError(w, err)
				return
			}

			for _, role := range roles {
				for _, userRole := range userPayload.Roles {
					if role == userRole {
						next.ServeHTTP(w, r)
						return
					}
				}
			}

			exc := NewSingleMessageException(
				EUNAUTHORIZED,
				"RequireRoles",
				"Not Authorized",
				errors.New("trying to access endpoint without appropiate role"),
			)
			SendError(w, exc)
		})
	}
}
