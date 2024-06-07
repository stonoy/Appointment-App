package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stonoy/Appointment-App/internal/database"
)

type apiConfig struct {
	fileServerHit int
	jwt_secret    string
	DB            *database.Queries
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

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		log.Println("No jwt secret set")
	}

	// configure apiConfig
	apiCfg := &apiConfig{
		fileServerHit: 0,
		jwt_secret:    jwt_secret,
	}

	db_conn := os.Getenv("DB_CONN")
	if port == "" {
		log.Printf("No Database set : %v", db_conn)
	} else {
		db, err := sql.Open("postgres", db_conn)
		if err != nil {
			log.Fatalf("can not open db connection : %v", err)
		}

		apiCfg.DB = database.New(db)

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

	// user
	apiRouter.Post("/register", apiCfg.register)
	apiRouter.Post("/login", apiCfg.login)

	// patient
	apiRouter.Post("/createpatient", apiCfg.onlyForAuthinticatedUser(apiCfg.createPatient))
	apiRouter.Post("/createappointment", apiCfg.onlyForAuthinticatedUser(apiCfg.createAppointment))
	apiRouter.Get("/getavailabilities", apiCfg.getAvailability)
	apiRouter.Get("/getappointments", apiCfg.onlyForAuthinticatedUser(apiCfg.getAppointments))
	apiRouter.Delete("/deleteappointments/{appointmentId}", apiCfg.onlyForAuthinticatedUser(apiCfg.DeleteAppointment))

	// doctor
	apiRouter.Post("/createavailability", apiCfg.onlyForDoctor(apiCfg.createAvailability))
	apiRouter.Get("/getavailabilitiesdoctor", apiCfg.onlyForDoctor(apiCfg.getAvailabilityDoctor))

	// admin
	apiRouter.Post("/createdoctor", apiCfg.onlyForAdmin(apiCfg.createDoctors))
	apiRouter.Get("/getalldoctors", apiCfg.onlyForAdmin(apiCfg.getAllDoctors))

	// mount sub router over main router
	mainRouter.Mount("/api/v1", apiRouter)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	log.Printf("Server is listening on port : %v", port)

	log.Fatal(server.ListenAndServe())
}
