package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"technical_test_Go/backend/api/routes"

	"technical_test_Go/backend/storage"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func main() {
	flag.Parse()

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	}))

	dbConnection, dbclose := storage.DatabaseConnection()

	// RESTy routes for "articles" resource
	router.Route("/load", func(r chi.Router) {
		r.Get("/", routes.DataLoader(dbConnection))
	})
	router.Route("/load/{date}", func(r chi.Router) {
		r.Get("/", routes.DataLoader(dbConnection))
	})
	router.Route("/allbuyers", func(r chi.Router) {
		r.Get("/", routes.GetAllBuyers(dbConnection))
	})
	router.Route("/buyers/{id}", func(r chi.Router) {
		r.Post("/", routes.GetDataBuyerbyID(dbConnection))
	})
	router.Route("/allproducts", func(r chi.Router) {
		r.Get("/", routes.GetAllProducts(dbConnection))
	})
	router.Route("/alltransactions", func(r chi.Router) {
		r.Get("/", routes.GetAllTransactions(dbConnection))
	})

	defer dbclose()

	fmt.Println("the service is running on port 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("Error listening and serving")
	}
}
