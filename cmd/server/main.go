package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Настройка маршрутов
	mux := http.NewServeMux()

	// Статические файлы
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Основные маршруты
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)

	// Запуск сервера
	log.Println("🚀 Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("❌ Ошибка:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>PC Builder</title>
		<style>
			body { font-family: Arial; margin: 40px; }
			h1 { color: #007acc; }
		</style>
	</head>
	<body>
		<h1>✅ PC Builder App работает!</h1>
		<p>Сервер запущен успешно!</p>
		<a href="/health">Проверить здоровье сервера</a>
	</body>
	</html>
	`)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"status": "ok"}`)
}
