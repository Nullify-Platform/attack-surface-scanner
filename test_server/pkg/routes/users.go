package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nullify-platform/logger/pkg/logger"
)

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func UsersRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		var newUser NewUser
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Info(
			"received new user request",
			logger.Any("newUser", newUser),
		)

		w.WriteHeader(http.StatusCreated)
	})

	r.Get("/{username}", func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		user := User{
			ID:       "123",
			Username: username,
			Email:    "myemail@example.com",
		}

		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return r
}
