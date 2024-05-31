package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	fileServerHit int
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error in loading .env... defailt configuration activated")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No Port set")
	}

	// configure apiConfig
	apiCfg := &apiConfig{
		fileServerHit: 0,
	}

	db_conn := os.Getenv("DB_CONN")
	if port == "" {
		log.Printf("No Database set : %v", db_conn)
	} else {
		// ...

		log.Println("Database Connected")
	}

	// main handler
	mainRouter := chi.NewRouter()

	// make it cors enable
	mainRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// sub router
	apiRouter := chi.NewRouter()

	// check health
	apiRouter.Get("/checkhealth", apiCfg.checkOk)
	apiRouter.Get("/checkerror", apiCfg.checkError)

	// mount sub router over main router
	mainRouter.Mount("/api/v1", apiRouter)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	log.Printf("Server is listening on port : %v", port)

	log.Fatal(server.ListenAndServe())
}
