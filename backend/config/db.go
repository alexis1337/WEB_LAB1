package db

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func LoadEnvVariables() error {
	err := godotenv.Load("config/databasesql.env")
	if err != nil {
		return fmt.Errorf("не удалось загрузить .env файл: %v", err)
	}
	return nil
}

func Connect() error {
	dbUser := GetEnvVariable("DB_USER")
	dbPassword := GetEnvVariable("DB_PASSWORD")
	dbName := GetEnvVariable("DB_NAME")
	dbSslMode := GetEnvVariable("DB_SSLMODE")

	dbHost := GetEnvVariable("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
		dbSslMode,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}

	connStr = fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
		dbSslMode,
	)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных")

	err = EnsureNewsTableExists(DB)
	if err != nil {
		return fmt.Errorf("ошибка проверки/создания таблицы: %v", err)
	}

	return nil
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}

func EnsureNewsTableExists(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS news (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы news: %v", err)
	}

	fmt.Println("Таблица news проверена или успешно создана")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Соединение с базой данных закрыто")
	}
}

func SetupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		CloseDB()
		os.Exit(0)
	}()
}
