package main

import (
	"fmt"
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
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/db-check", dbCheckHandler)

	// Страницы компонентов
	mux.HandleFunc("/components", handlers.ComponentsPage)                // Каталог всех компонентов
	mux.HandleFunc("/components/cpus", handlers.CPUsPage)                 // Страница процессоров
	mux.HandleFunc("/components/gpus", handlers.GPUsPage)                 // Страница видеокарт
	mux.HandleFunc("/components/motherboards", handlers.MotherboardsPage) // Страница материнских плат
	mux.HandleFunc("/components/rams", handlers.RAMsPage)                 // Страница оперативной памяти

	// API маршруты
	mux.HandleFunc("/api/cpus", handlers.GetCPUs)                 // API процессоров
	mux.HandleFunc("/api/gpus", handlers.GetGPUs)                 // API видеокарт
	mux.HandleFunc("/api/motherboards", handlers.GetMotherboards) // API материнских плат
	mux.HandleFunc("/api/rams", handlers.GetRAMs)                 // API оперативной памяти

	// Запуск сервера
	log.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Println("📊 База данных подключена")
	log.Println("🖥️  Главная: http://localhost:8080")
	log.Println("📦 Каталог компонентов: http://localhost:8080/components")
	log.Println("⚙️  API процессоров: http://localhost:8080/api/cpus")
	err = http.ListenAndServe(":8080", mux)
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
			.status { padding: 10px; border-radius: 5px; margin: 10px 0; }
			.success { background: #d4edda; color: #155724; }
			.nav { margin: 20px 0; }
			.btn { display: inline-block; background: #007acc; color: white; padding: 10px 20px; text-decoration: none; border-radius: 4px; margin: 5px; }
			.menu { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 10px; margin: 20px 0; }
			.menu-item { border: 1px solid #ddd; padding: 15px; border-radius: 5px; text-align: center; }
		</style>
	</head>
	<body>
		<h1>✅ PC Builder App работает!</h1>
		<div class="status success">📊 База данных подключена и настроена</div>
		
		<div class="nav">
			<a href="/health" class="btn">Проверить здоровье сервера</a>
			<a href="/db-check" class="btn">Проверить базу данных</a>
			<a href="/components" class="btn">📦 Все компоненты</a>
		</div>
		
		<h2>🔧 Меню компонентов:</h2>
		<div class="menu">
			<div class="menu-item">
				<h3>🖥️ Процессоры</h3>
				<p>Intel, AMD и другие</p>
				<a href="/components/cpus">Перейти →</a>
			</div>
			<div class="menu-item">
				<h3>🎮 Видеокарты</h3>
				<p>NVIDIA, AMD Radeon</p>
				<a href="/components/gpus">Перейти →</a>
			</div>
			<div class="menu-item">
				<h3>🔌 Материнские платы</h3>
				<p>ASUS, Gigabyte, MSI</p>
				<a href="/components/motherboards">Перейти →</a>
			</div>
			<div class="menu-item">
				<h3>💾 Оперативная память</h3>
				<p>DDR4, DDR5</p>
				<a href="/components/rams">Перейти →</a>
			</div>
		</div>
		
		<h3>Что уже работает:</h3>
		<ul>
			<li>✅ Веб-сервер на Go</li>
			<li>✅ PostgreSQL база данных</li>
			<li>✅ Автоматические миграции</li>
			<li>✅ Тестовые данные (процессоры, видеокарты и др.)</li>
			<li>✅ API для получения компонентов</li>
			<li>✅ Веб-интерфейс для просмотра</li>
		</ul>
	</body>
	</html>
	`)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"status": "ok", "service": "pc-builder"}`)
}

func dbCheckHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	err := database.DB.QueryRow("SELECT 'База данных работает! Таблицы созданы автоматически.'").Scan(&result)
	if err != nil {
		http.Error(w, `{"status": "error", "message": "Ошибка БД"}`, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status": "ok", "message": "%s"}`, result)
}
