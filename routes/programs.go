package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zehuxx/python-code-api/handlers"
	"github.com/zehuxx/python-code-api/middlewares"
)

//Programs router
func Programs() http.Handler {
	router := chi.NewRouter()

	router.Route("/programs", func(r chi.Router) {
		r.Get("/", handlers.GetPrograms)
		r.Post("/", handlers.SaveProgram)
		r.Route("/{uid}", func(r chi.Router) {
			r.Use(middlewares.GetUid)
			r.Get("/", handlers.GetProgramByUid)
			r.Put("/", handlers.UpdateProgram)
			r.Delete("/", handlers.DeleteProgram)
		})

	})

	return router
}
