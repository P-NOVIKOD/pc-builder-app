package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Init инициализирует подключение к базе данных и запускает миграции
func Init() error {
	connStr := "user=postgres password=2287 dbname=pcbuilder host=localhost port=5432 sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	// Настройки пула соединений
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Проверяем подключение
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("не удалось подключиться к БД. Убедись что PostgreSQL запущен: %w", err)
	}

	log.Println("✅ Подключение к PostgreSQL установлено")

	// Запускаем миграции
	err = runMigrations()
	if err != nil {
		return fmt.Errorf("ошибка миграций: %w", err)
	}

	log.Println("✅ Миграции базы данных выполнены")
	return nil
}

// runMigrations выполняет все SQL миграции
func runMigrations() error {
	// 1. Создаем таблицы
	tablesSQL, err := os.ReadFile("migrations/sql/001_create_tables.sql")
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(tablesSQL))
	if err != nil {
		return err
	}

	// 2. Создаем триггеры
	triggersSQL, err := os.ReadFile("migrations/sql/003_create_triggers.sql")
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(triggersSQL))
	if err != nil {
		return err
	}

	log.Println("✅ Таблицы и триггеры созданы")
	return nil
}

// Close закрывает подключение к базе данных
func Close() {
	if DB != nil {
		DB.Close()
	}
}
