package main

import (
	"log"
	"net/http"
	"os"
	"waste_management/config"
	"waste_management/controller"
	"waste_management/model"
	handler "waste_management/view/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем доступ с любого источника
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Указываем методы, которые разрешены
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// Указываем разрешённые заголовки
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Если это OPTIONS-запрос, завершаем его без передачи дальше
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Передаём управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Read .env file and create connection config

	err := godotenv.Load("config/.env")

	if err != nil {
		log.Println("Error while reading dotenv: ", err.Error())
	}

	conf := &config.MongoDbConfig{
		Uri: os.Getenv("MONGODB_URI"),
		Database: os.Getenv("MONGODB_DATABASE"),
	}

	// Init modules

	connection := model.NewMongoDBConnection(conf)

	repository := model.NewRepository(connection)

	controller := controller.NewController(repository)

	technologyHandler := handler.NewTechnologyHandler(controller)
	producerHandler := handler.NewProducerHandler(controller)
	fkkoHandler := handler.NewFkkoHandler(controller)
	okpdHandler := handler.NewOkpdHandler(controller)

	// Register paths of API

	router := mux.NewRouter()

	router.HandleFunc("/technology", technologyHandler.GetTechnology).Methods(http.MethodGet)
	router.HandleFunc("/technology", technologyHandler.PostTechnology).Methods(http.MethodPost)
	router.HandleFunc("/technologies", technologyHandler.GetTechnologies).Methods(http.MethodGet)

	router.HandleFunc("/producer", producerHandler.PostProducer).Methods(http.MethodPost)
	router.HandleFunc("/producers", producerHandler.GetProducers).Methods(http.MethodGet)

	router.HandleFunc("/fkkos", fkkoHandler.GetFkkos).Methods(http.MethodGet)
	router.HandleFunc("/okpds", okpdHandler.GetOkpds).Methods(http.MethodGet)

	routerWithCorse := corsMiddleware(router)

	// Start server
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", routerWithCorse); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}