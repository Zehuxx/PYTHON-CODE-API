package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zehuxx/python-code-api/db"
	"github.com/zehuxx/python-code-api/handlers"
	"github.com/zehuxx/python-code-api/middlewares"
)

//Programs router
func Programs(dcl *db.DgraphInstance) http.Handler {
	router := chi.NewRouter()

	router.Route("/programs", func(r chi.Router) {
		r.Get("/", handlers.GetPrograms(dcl))
		r.Post("/", handlers.SaveProgram(dcl))
		r.Route("/{uid}", func(r chi.Router) {
			r.Use(middlewares.GetUid)
			r.Get("/", handlers.GetProgramByUid(dcl))
			r.Put("/", handlers.UpdateProgram(dcl))
			r.Delete("/", handlers.DeleteProgram(dcl))
		})

	})

	return router
}
