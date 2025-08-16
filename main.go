package main

import (
	"Notes_TodoList/internal/handler"
    "Notes_TodoList/internal/repository/postgres"
    "Notes_TodoList/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=Noname0212 dbname=HomeWork.16 sslmode=disable")
	if err != nil {
		log.Fatalln("Ошибка подключения к базе данных:", err)
	}

	repo := postgres.NewCategoryPostgres(db)
	uc := usecase.NewCategoryUsecase(repo)

	r := gin.Default()
	handler.NewCategoryHandler(r, &uc)

	if err := r.Run(":6666"); err != nil {
		log.Fatalln("Ошибка запуска сервера:", err)
	}
}
