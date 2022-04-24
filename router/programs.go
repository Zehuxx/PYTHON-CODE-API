package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/zehuxx/python-code-api/db"
	"github.com/zehuxx/python-code-api/routes"
)

//StartServer start up server
func StartServer() {
	router := chi.NewRouter()

	//middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.AllowAll().Handler)

	//get dgraph client
	dgraphcl, closeConnection := db.GetDgraphClient()
	defer closeConnection()

	//program routes
	router.Mount("/api", routes.Programs(dgraphcl))

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	fmt.Print("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
