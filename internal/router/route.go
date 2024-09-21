package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jexroid/gopi/internal/crud"

	"github.com/jexroid/gopi/internal/database"
	"github.com/jexroid/gopi/internal/handler"
)

func Handler(route *chi.Mux) {
	// global middleware
	route.Use(middleware.StripSlashes)
	var DB = database.Init()
	AuthDB := handler.InitDB(DB)
	CrudDB := crud.InitDB(DB)

	route.Route("/auth", func(router chi.Router) {
		router.Post("/signup", AuthDB.Signup)
		router.Post("/signin", AuthDB.Signin)
		router.Post("/validate", AuthDB.Validate)
		router.Get("/logout", handler.Logout)
	})

	route.Route("/cars", func(router chi.Router) {
		router.Post("/", CrudDB.Create)         // Create a new car
		router.Get("/{uuid}", CrudDB.Read)      // Read a car by UUID
		router.Get("/", CrudDB.All)             // Read a car by UUID
		router.Put("/{uuid}", CrudDB.Update)    // Update a car by UUID
		router.Delete("/{uuid}", CrudDB.Delete) // Delete a car by UUID
	})
}
