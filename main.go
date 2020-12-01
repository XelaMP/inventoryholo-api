package main

import (
	"fmt"
	"github.com/XelaMP/inventoryholo-api/constants"
	"github.com/XelaMP/inventoryholo-api/db"
	"github.com/XelaMP/inventoryholo-api/helper"
	"github.com/XelaMP/inventoryholo-api/middleware"
	routes "github.com/XelaMP/inventoryholo-api/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main ()  {
	r := mux.NewRouter()

	db.DB = helper.Get()

	r.HandleFunc("/", indexRouter)
	r.HandleFunc("/api/login", middleware.Login)

	s := r.PathPrefix("/api").Subrouter()
	routes.Routes(s)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{
			"http://localhost:4200",
			"http://192.241.159.224",
			"http://resultados.holosalud.org",
			"https://resultados.holosalud.org",
		},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "x-token"},
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = constants.PORT //localhost
	}

	handler := c.Handler(r)

	fmt.Println("Server online!")

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func indexRouter(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome api inventory holo!")
}

