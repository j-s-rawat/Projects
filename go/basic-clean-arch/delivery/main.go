// delivery/main.go
package main

import (
	"clean_arch/delivery/http"
	"clean_arch/repository"
	"clean_arch/usecase"
	h "net/http"
)

func main() {
	userRepository := repository.NewUserRepositoryImpl()
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := http.NewUserHandler(userUseCase)

	router := http.NewRouter(userHandler)

	// Start the HTTP server
	h.ListenAndServe(":8080", router)
}
