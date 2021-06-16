package common

import (
	"context"
	"errors"
	"net/http"
)

type key struct{}

func UserPayloadFromContext(op string, r *http.Request) (*UserPayload, error) {
	userPayload, ok := r.Context().Value(key{}).(*UserPayload)
	if !ok {
		return nil, &Exception{
			Op:  op,
			Err: errors.New("missing user payload"),
		}
	}

	return userPayload, nil
}

func UserPayloadToContext(userPayload *UserPayload, r *http.Request) context.Context {
	return context.WithValue(r.Context(), key{}, userPayload)
}
