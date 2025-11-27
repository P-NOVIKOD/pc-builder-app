package handlers

import (
	"encoding/json"
	"net/http"
)

// HandleLogin обрабатывает вход пользователя
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Пока просто заглушка
	response := map[string]interface{}{
		"success": true,
		"user_id": 1,
		"message": "Вход выполнен",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetComponents возвращает все компоненты из БД
func GetComponents(w http.ResponseWriter, r *http.Request) {
	// Пока заглушка - вернем тестовые данные
	components := []map[string]interface{}{
		{
			"id":         1,
			"type":       "cpu",
			"vendor":     "AMD",
			"model":      "Ryzen 7 7800X3D",
			"price":      1150.00,
			"core_count": 8,
			"tdp":        120,
		},
		{
			"id":        2,
			"type":      "gpu",
			"vendor":    "NVIDIA",
			"model":     "GeForce RTX 3070 Ti",
			"price":     2871.00,
			"vram_size": 8,
			"tdp":       290,
		},
		{
			"id":      3,
			"type":    "motherboard",
			"vendor":  "ASUS",
			"model":   "TUF Gaming X870-Plus WiFi",
			"price":   1260.10,
			"chipset": "X870",
		},
		{
			"id":          4,
			"type":        "ram",
			"vendor":      "KINGBANK",
			"model":       "KRRB 2x16ГБ",
			"price":       1125.00,
			"capacity_gb": 32,
			"speed_mhz":   6000,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(components)
}

// HandleBuilds управляет сборками (создание, получение)
func HandleBuilds(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBuilds(w, r)
	case "POST":
		createBuild(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// getBuilds возвращает список сборок пользователя
func getBuilds(w http.ResponseWriter, r *http.Request) {
	// Пока заглушка
	builds := []map[string]interface{}{
		{
			"id":    1,
			"name":  "Игровой ПК",
			"price": 5000.00,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(builds)
}

// createBuild создает новую сборку
func createBuild(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name       string                 `json:"name"`
		Components map[string]interface{} `json:"components"`
	}

	// Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	// Пока просто возвращаем успех
	response := map[string]interface{}{
		"success": true,
		"message": "Сборка сохранена",
		"build": map[string]interface{}{
			"name":  request.Name,
			"price": calculateTotalPrice(request.Components),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// calculateTotalPrice вычисляет общую стоимость сборки
func calculateTotalPrice(components map[string]interface{}) float64 {
	total := 0.0
	for _, comp := range components {
		if component, ok := comp.(map[string]interface{}); ok {
			if price, ok := component["price"].(float64); ok {
				total += price
			}
		}
	}
	return total
}
