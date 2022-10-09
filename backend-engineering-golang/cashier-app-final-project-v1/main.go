package main

import (
	"a21hc3NpZ25tZW50/api"
	"a21hc3NpZ25tZW50/db"
	repo "a21hc3NpZ25tZW50/repository"
)

func main() {
	db := db.NewJsonDB()
	usersRepo := repo.NewUserRepository(db)
	sessionsRepo := repo.NewSessionsRepository(db)
	productsRepo := repo.NewProductRepository(db)
	cartsRepo := repo.NewCartRepository(db)

	mainAPI := api.NewAPI(usersRepo, sessionsRepo, productsRepo, cartsRepo)
	mainAPI.Start()
}
