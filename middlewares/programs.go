package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zehuxx/python-code-api/helpers"
)

//GetUid test to send information through the context.
func GetUid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "uid")
		if uid == "" {
			helpers.JSONError(w, helpers.ErrorResponse{Msg: "uid is required."}, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", uid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

/*
You could have more middleware to validate the uid against dgraph to see if it exists,
 before modifying or deleting.
*/
