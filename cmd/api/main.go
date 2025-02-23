package main

import "fmt"
import "net/http"
import "github.com/go-chi/chi/v5"
import "github.com/go-chi/chi/v5/middleware"

import "github.com/muse/users"

func run_web_server() {
	fmt.Println("Starting server")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi!"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", users.CreateUserView)
	})

	server := &http.Server{Addr: ":3000", Handler: r}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func main() {
	run_web_server()
}
