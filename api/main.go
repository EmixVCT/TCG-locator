package main

import (
	"log"
	"net/http"
	"tcglocator/controller"
	"tcglocator/utils"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func init() {
	err := utils.LoadConfig("config/settings.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCG-Locator configuration file loaded, version:", utils.GetConfig().Version)

	utils.InitCrawl()
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json", "text/xml"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://localhost", "http://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Use(httprate.LimitByIP(10, 1*time.Minute))
		r.Use(middleware.Heartbeat("/"))

		r.Get("/locators", controller.Locator)
	})

	r.NotFound(controller.NotFound)
	r.MethodNotAllowed(controller.NotValid)

	// Require Authentication
	// r.Group(func(r chi.Router) {
	//     r.Use(AuthMiddleware)
	//     r.Post("/manage", CreateAsset)
	// })

	log.Println("Server running on port 3050")
	http.ListenAndServe(":3050", r)
}
