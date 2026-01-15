package main

import (
	"fmt"
	"net/http"

	_ "github.com/giordanGarci/api-tenants/docs"
	"github.com/giordanGarci/api-tenants/handlers"
	"github.com/giordanGarci/api-tenants/interceptors"
	"github.com/giordanGarci/api-tenants/repository"
	"github.com/giordanGarci/api-tenants/services"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Bots API
// @version 1.0
// @description API para gerenciamento de bots
// @host localhost:8080
// @BasePath /
func main() {
	// Dependency Injection: criamos as dependências de baixo para cima
	repo := repository.NewInMemoryBotRepository()
	service := services.NewService(repo)
	botHandler := handlers.NewBotHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)

	middlewareGetBots := interceptors.EnsureRole("admin", "dev", "user")
	mux.Handle("/bots", middlewareGetBots(http.HandlerFunc(botHandler.GetAllBotsHandler)))
	mux.Handle("/bot", middlewareGetBots(http.HandlerFunc(botHandler.GetBotByIDHandler)))

	middlewareCreateBot := interceptors.EnsureRole("admin", "dev")
	createChain := middlewareCreateBot(http.HandlerFunc(botHandler.CreateBotHandler))
	mux.Handle("/bot/create", createChain)

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("server listening on :8080")
	middlewaresStack := interceptors.AuthMiddleware(mux)
	if err := http.ListenAndServe(":8080", middlewaresStack); err != nil {
		panic(err)
	}

}
