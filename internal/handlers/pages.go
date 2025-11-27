package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// LoginPage отображает страницу входа
func LoginPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html")
}

// BuilderPage отображает страницу конструктора
func BuilderPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "builder.html")
}

// AdminPage отображает страницу админки
func AdminPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "admin.html")
}

// renderTemplate загружает и отображает HTML шаблон
func renderTemplate(w http.ResponseWriter, templateName string) {
	// Путь к файлу шаблона
	filePath := filepath.Join("templates", "pages", templateName)

	// Проверяем существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "Страница не найдена: "+filePath, http.StatusNotFound)
		return
	}

	// Читаем файл
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Страница не найдена", http.StatusNotFound)
		return
	}

	// Устанавливаем заголовок и выводим HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, string(content))
}
