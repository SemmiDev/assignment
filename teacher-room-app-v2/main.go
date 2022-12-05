package main

import (
	"embed"

	"a21hc3NpZ25tZW50/app/controller"
	"a21hc3NpZ25tZW50/app/model"
	repo "a21hc3NpZ25tZW50/app/repository"
	"a21hc3NpZ25tZW50/config"

	_ "github.com/jackc/pgx/v4/stdlib"
)

//go:embed app/view/*
var Resources embed.FS

func main() {
	// os.Setenv("DATABASE_URL", "postgres://root:secret@localhost:5432/assignments") // Hapus jika akan melakukan deploy ke fly.io

	db := config.NewDB()
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.Teacher{})
	teacherHandle := repo.NewTeacherRepo(conn)

	mainAPI := controller.NewAPI(teacherHandle, Resources)
	mainAPI.Start()
}

func FlyURL() string {
	// TODO: replace this with your fly.io domain url, example: return "https://aditira-2022.fly.dev"

	return "https://sammi-aldhi-yanto-be2808835.fly.dev" // TODO: replace this
}
