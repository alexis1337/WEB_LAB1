package main

import (
	"fmt"
	"log"
	db "news_app/config"
	"news_app/controllers"
	"news_app/repository"
	"news_app/routes"
	"news_app/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := db.LoadEnvVariables(); err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	validateEnvVars()

	if err := db.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Printf("Ошибка при закрытии соединения с БД: %v", err)
		}
	}()

	if err := migrateDatabase(); err != nil {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}

	startServer()
}

func validateEnvVars() {
	requiredEnvVars := []string{"DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSLMODE"}
	for _, key := range requiredEnvVars {
		if db.GetEnvVariable(key) == "" {
			log.Fatalf("Отсутствует обязательная переменная окружения: %s", key)
		}
	}
}

func migrateDatabase() error {
	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
			db.GetEnvVariable("DB_USER"),
			db.GetEnvVariable("DB_PASSWORD"),
			db.GetEnvVariable("DB_HOST"),
			db.GetEnvVariable("DB_NAME"),
			db.GetEnvVariable("DB_SSLMODE"),
		),
	)
	if err != nil {
		return fmt.Errorf("ошибка создания миграции: %v", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Printf("Ошибка применения миграции: %v", err)
		return fmt.Errorf("ошибка применения миграции: %v", err)
	}

	fmt.Println("Миграции выполнены успешно")
	return nil
}

func startServer() {
	repo := repository.NewNewsRepository(db.DB)
	svc := service.NewNewsService(repo)
	ctrl := controllers.NewNewsController(svc)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(r, ctrl)

	port := ":8080"
	fmt.Printf("Сервер запущен на порту %s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
