package main

import (
	"avito-test-pvz/internal/config"
	"avito-test-pvz/internal/handler"
	"avito-test-pvz/internal/repository"
	"avito-test-pvz/internal/server"
	"avito-test-pvz/internal/service"
	"avito-test-pvz/internal/storage"
	"fmt"
	"log"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	store, err := storage.New(cfg.StoragePath)
	if err != nil {
		log.Fatal("storage error:", err)
	}

	userRepo := repository.NewUserRepository(store)

	userService := service.New(userRepo)
	h := handler.New(userService)

	s := server.New(cfg, store)
	s.Run(h)

}
