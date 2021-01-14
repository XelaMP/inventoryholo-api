package main

import (
	"flag"
	"fmt"
	"github.com/XelaMP/inventoryholo-api/constants"
	"github.com/XelaMP/inventoryholo-api/db"
	"github.com/XelaMP/inventoryholo-api/helper"
	"github.com/XelaMP/inventoryholo-api/middleware"
	routes "github.com/XelaMP/inventoryholo-api/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/gonutz/w32"
	"log"
	"net/http"
	"os"
)

func main() {
	f := flag.Bool("execTerminal", false, "to exec cmd")
	flag.Parse()
	if !*f {
		hideConsole()
	}
	api()

}

	func api(){
		r := mux.NewRouter()

		db.DB = helper.Get()

		r.HandleFunc("/", indexRouter)
		r.HandleFunc("/api/login", middleware.Login)

		s := r.PathPrefix("/api").Subrouter()
		routes.Routes(s)

		c := cors.New(cors.Options{
			AllowedOrigins: []string{
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

func hideConsole(){
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}
	}
}
