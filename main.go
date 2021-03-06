package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/adrianhosman/structural-design-go/config"
	marveldal "github.com/adrianhosman/structural-design-go/dal/api/marvel"
	cachedal "github.com/adrianhosman/structural-design-go/dal/cache"
	repodal "github.com/adrianhosman/structural-design-go/dal/repo"
	"github.com/adrianhosman/structural-design-go/handler"
	"github.com/adrianhosman/structural-design-go/resources"
	"github.com/adrianhosman/structural-design-go/usecase"
	cron "github.com/robfig/cron/v3"
)

func main() {
	// Init resources
	cfg := config.GetConfig("./config/")
	marvel := resources.InitMarvelClient(cfg)
	cache := resources.InitCache(cfg)

	// Init layers
	marvelDAL := marveldal.New(cfg, marvel)
	cacheDAL := cachedal.New(cache)
	repoDal := repodal.New()
	usecaseLayer := usecase.New(cfg, marvelDAL, cacheDAL, repoDal)
	handlerLayer := handler.New(usecaseLayer)

	// initNecessaryData(usecaseLayer)
	// initCron(usecaseLayer)
	initHTTP(cfg, handlerLayer)
}

// Init necessary data
func initNecessaryData(usecaseLayer usecase.Usecase) {
	go func() {
		defer func() { // recover go routine in case of panic
			if r := recover(); r != nil {
				log.Printf("error = %v", fmt.Errorf("%v", r))
			}
		}()
		usecaseLayer.SaveCharacters()
	}()
}

// Init cron
func initCron(usecaseLayer usecase.Usecase) {
	cronHandler := cron.New()
	cronHandler.AddFunc("@hourly", func() { usecaseLayer.SaveCharacters() }) // update marvel characters cache hourly (in case of new characters)
}

// Init routes
func initHTTP(cfg *config.Config, handlerLayer *handler.Handler) {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet)
	// r.HandleFunc("/characters/{id}", handlerLayer.GetCharacterByID).Methods(http.MethodGet)
	// r.HandleFunc("/characters", handlerLayer.GetAllCharacterIDs).Methods(http.MethodGet)
	r.HandleFunc("/invoice/calculation/{business_id}", handlerLayer.CalculateInvoice).Methods(http.MethodGet)

	// Start server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type"},
	})
	handler := c.Handler(r)
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server started at %s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
