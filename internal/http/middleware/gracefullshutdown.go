package middleware

import (
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/internal/services"
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func GracefulShutdownMiddlleware(taskservice services.TaskService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if domain.IsDraining() {

				taskBytes, err := io.ReadAll(r.Body)
				if err != nil {
					render.Status(r, http.StatusInternalServerError)
					return
				}

				defer r.Body.Close()

				task := domain.Task{
					ID:     uuid.NewString(),
					Date:   string(taskBytes),
					Path:   r.URL.Path,
					Method: r.Method,
				}

				if err := taskservice.SaveTask(r.Context(), task); err != nil {
					render.Status(r, http.StatusInternalServerError)
					return
				}

				render.Status(r, http.StatusServiceUnavailable)
				return

			}

			next.ServeHTTP(w, r)
		})
	}

}
