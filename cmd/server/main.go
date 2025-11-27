package main

import (
	"log"
	"net/http"
	"pc-builder/internal/database"
	"pc-builder/internal/handlers"
)

func main() {
	// Инициализируем базу данных
	err := database.Init()
	if err != nil {
		log.Fatal("❌ Ошибка инициализации БД:", err)
	}
	defer database.Close()

	// Настройка маршрутов
	mux := http.NewServeMux()

	// Статические файлы
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Основные маршруты
	mux.HandleFunc("/", handlers.LoginPage)          // Главная = вход
	mux.HandleFunc("/login", handlers.LoginPage)     // Страница входа
	mux.HandleFunc("/builder", handlers.BuilderPage) // Конструктор сборок
	mux.HandleFunc("/admin", handlers.AdminPage)     // Добавление компонентов

	// API маршруты
	mux.HandleFunc("/api/login", handlers.HandleLogin)        // Авторизация
	mux.HandleFunc("/api/components", handlers.GetComponents) // Все компоненты
	mux.HandleFunc("/api/builds", handlers.HandleBuilds)      // Управление сборками

	// Запуск сервера
	log.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Println("📊 База данных подключена")
	log.Println("🔐 Вход: http://localhost:8080")
	log.Println("🛠️  Конструктор: http://localhost:8080/builder")
	log.Println("⚙️  Админка: http://localhost:8080/admin")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("❌ Ошибка:", err)
	}
}
