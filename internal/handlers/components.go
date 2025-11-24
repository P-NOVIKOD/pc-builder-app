// CPUsPage возвращает HTML страницу с процессорами
func CPUsPage(w http.ResponseWriter, r *http.Request) {
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Процессоры - PC Builder</title>
        <style>
            body { font-family: Arial; margin: 40px; }
            .component { border: 1px solid #ddd; padding: 15px; margin: 10px 0; border-radius: 5px; }
            .cpu { border-left: 4px solid #007acc; }
            .price { color: #e44d26; font-weight: bold; }
            .nav { margin-bottom: 20px; }
        </style>
    </head>
    <body>
        <div class="nav">
            <a href="/">← На главную</a> | 
            <a href="/components">Все компоненты</a>
        </div>
        <h1>🖥️ Процессоры</h1>
        <div id="cpus-list">
            Загрузка...
        </div>

        <script>
            fetch('/api/cpus')
                .then(response => response.json())
                .then(cpus => {
                    const container = document.getElementById('cpus-list');
                    container.innerHTML = cpus.map(cpu => 
                        '<div class="component cpu">' +
                        '<h3>' + cpu.vendor + ' ' + cpu.model + '</h3>' +
                        '<p>Сокет: ' + cpu.socket + ' | Ядра: ' + cpu.core_count + ' | Потоки: ' + cpu.thread_count + '</p>' +
                        '<p>Частота: ' + cpu.clock_speed + ' GHz | TDP: ' + cpu.power_consumption + 'W</p>' +
                        '<p class="price">' + cpu.price + ' ₽</p>' +
                        '</div>'
                    ).join('');
                });
        </script>
    </body>
    </html>
    `
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

// GPUsPage возвращает HTML страницу с видеокартами
func GPUsPage(w http.ResponseWriter, r *http.Request) {
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Видеокарты - PC Builder</title>
        <style>
            body { font-family: Arial; margin: 40px; }
            .component { border: 1px solid #ddd; padding: 15px; margin: 10px 0; border-radius: 5px; }
            .gpu { border-left: 4px solid #28a745; }
            .price { color: #e44d26; font-weight: bold; }
            .nav { margin-bottom: 20px; }
        </style>
    </head>
    <body>
        <div class="nav">
            <a href="/">← На главную</a> | 
            <a href="/components">Все компоненты</a>
        </div>
        <h1>🎮 Видеокарты</h1>
        <div id="gpus-list">
            Загрузка...
        </div>

        <script>
            fetch('/api/gpus')
                .then(response => response.json())
                .then(gpus => {
                    const container = document.getElementById('gpus-list');
                    container.innerHTML = gpus.map(gpu => 
                        '<div class="component gpu">' +
                        '<h3>' + gpu.vendor + ' ' + gpu.model + '</h3>' +
                        '<p>VRAM: ' + gpu.vram_gb + ' GB | Тип памяти: ' + gpu.memory_type + '</p>' +
                        '<p>TDP: ' + gpu.power_consumption + 'W</p>' +
                        '<p class="price">' + gpu.price + ' ₽</p>' +
                        '</div>'
                    ).join('');
                });
        </script>
    </body>
    </html>
    `
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

// GetGPUs возвращает список видеокарт
func GetGPUs(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, vendor, model, vram_gb, memory_type, price, power_consumption FROM gpus ORDER BY price")
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка базы данных: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type GPU struct {
		ID               int     `json:"id"`
		Vendor           string  `json:"vendor"`
		Model            string  `json:"model"`
		VRAMGB           int     `json:"vram_gb"`
		MemoryType       string  `json:"memory_type"`
		Price            float64 `json:"price"`
		PowerConsumption int     `json:"power_consumption"`
	}

	var gpus []GPU
	for rows.Next() {
		var gpu GPU
		err := rows.Scan(&gpu.ID, &gpu.Vendor, &gpu.Model, &gpu.VRAMGB, &gpu.MemoryType, &gpu.Price, &gpu.PowerConsumption)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка чтения данных: %v", err), http.StatusInternalServerError)
			return
		}
		gpus = append(gpus, gpu)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gpus)
}

// Добавь аналогичные функции для материнских плат и оперативной памяти...
// MotherboardsPage, GetMotherboards, RAMsPage, GetRAMs