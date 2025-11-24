package database

import (
	"database/sql"
	"fmt"
	"log"
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
	// Миграция 1: Создание таблиц
	_, err := DB.Exec(`
        -- Создание таблицы пользователей
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(50) UNIQUE NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password_hash VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );

        -- Создание таблицы процессоров
        CREATE TABLE IF NOT EXISTS cpus (
            id SERIAL PRIMARY KEY,
            vendor VARCHAR(50) NOT NULL,
            model VARCHAR(100) NOT NULL,
            socket VARCHAR(50) NOT NULL,
            core_count INTEGER NOT NULL,
            thread_count INTEGER NOT NULL,
            clock_speed DECIMAL(4,2) NOT NULL,
            price DECIMAL(10,2) NOT NULL,
            power_consumption INTEGER NOT NULL
        );

        -- Создание таблицы материнских плат
        CREATE TABLE IF NOT EXISTS motherboards (
            id SERIAL PRIMARY KEY,
            vendor VARCHAR(50) NOT NULL,
            model VARCHAR(100) NOT NULL,
            socket VARCHAR(50) NOT NULL,
            chipset VARCHAR(50) NOT NULL,
            form_factor VARCHAR(20) NOT NULL,
            memory_slots INTEGER NOT NULL,
            price DECIMAL(10,2) NOT NULL
        );

        -- Создание таблицы видеокарт
        CREATE TABLE IF NOT EXISTS gpus (
            id SERIAL PRIMARY KEY,
            vendor VARCHAR(50) NOT NULL,
            model VARCHAR(100) NOT NULL,
            vram_gb INTEGER NOT NULL,
            memory_type VARCHAR(20) NOT NULL,
            price DECIMAL(10,2) NOT NULL,
            power_consumption INTEGER NOT NULL
        );

        -- Создание таблицы оперативной памяти
        CREATE TABLE IF NOT EXISTS rams (
            id SERIAL PRIMARY KEY,
            vendor VARCHAR(50) NOT NULL,
            model VARCHAR(100) NOT NULL,
            type VARCHAR(10) NOT NULL,
            speed_mhz INTEGER NOT NULL,
            capacity_gb INTEGER NOT NULL,
            price DECIMAL(10,2) NOT NULL
        );

        -- Создание таблицы сборок
        CREATE TABLE IF NOT EXISTS builds (
            id SERIAL PRIMARY KEY,
            user_id INTEGER REFERENCES users(id),
            name VARCHAR(100) NOT NULL,
            total_price DECIMAL(10,2) DEFAULT 0,
            total_power_consumption INTEGER DEFAULT 0,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );

        -- Создание таблицы компонентов в сборках
        CREATE TABLE IF NOT EXISTS build_components (
            id SERIAL PRIMARY KEY,
            build_id INTEGER REFERENCES builds(id) ON DELETE CASCADE,
            component_type VARCHAR(20) NOT NULL,
            component_id INTEGER NOT NULL
        );
    `)
	if err != nil {
		return err
	}

	// Миграция 2: Добавление тестовых данных
	err = insertTestData()
	if err != nil {
		return err
	}

	return nil
}

// insertTestData добавляет тестовые данные
func insertTestData() error {
	// Проверяем есть ли уже данные в CPU
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM cpus").Scan(&count)
	if err != nil {
		return err
	}

	// Если данные уже есть - пропускаем
	if count > 0 {
		log.Println("✅ Тестовые данные уже существуют")
		return nil
	}

	log.Println("📝 Добавляем тестовые данные...")

	// Добавляем процессоры
	_, err = DB.Exec(`
        INSERT INTO cpus (vendor, model, socket, core_count, thread_count, clock_speed, price, power_consumption) VALUES
        ('Intel', 'Core i5-13600K', 'LGA1700', 14, 20, 3.50, 320.00, 125),
        ('AMD', 'Ryzen 7 7800X3D', 'AM5', 8, 16, 4.20, 450.00, 120),
        ('Intel', 'Core i7-14700K', 'LGA1700', 20, 28, 3.40, 420.00, 125),
        ('AMD', 'Ryzen 5 7600X', 'AM5', 6, 12, 4.70, 250.00, 105);
    `)
	if err != nil {
		return err
	}

	// Добавляем материнские платы
	_, err = DB.Exec(`
        INSERT INTO motherboards (vendor, model, socket, chipset, form_factor, memory_slots, price) VALUES
        ('ASUS', 'ROG STRIX B650-A GAMING WIFI', 'AM5', 'B650', 'ATX', 4, 280.00),
        ('Gigabyte', 'Z790 AORUS ELITE AX', 'LGA1700', 'Z790', 'ATX', 4, 320.00),
        ('MSI', 'B760 GAMING PLUS WIFI', 'LGA1700', 'B760', 'ATX', 4, 190.00),
        ('ASRock', 'B650E STEEL LEGEND', 'AM5', 'B650', 'ATX', 4, 260.00);
    `)
	if err != nil {
		return err
	}

	// Добавляем видеокарты
	_, err = DB.Exec(`
        INSERT INTO gpus (vendor, model, vram_gb, memory_type, price, power_consumption) VALUES
        ('NVIDIA', 'GeForce RTX 4070', 12, 'GDDR6X', 600.00, 200),
        ('AMD', 'Radeon RX 7800 XT', 16, 'GDDR6', 550.00, 263),
        ('NVIDIA', 'GeForce RTX 4060 Ti', 8, 'GDDR6', 450.00, 165),
        ('AMD', 'Radeon RX 7700 XT', 12, 'GDDR6', 450.00, 245);
    `)
	if err != nil {
		return err
	}

	// Добавляем оперативную память
	_, err = DB.Exec(`
        INSERT INTO rams (vendor, model, type, speed_mhz, capacity_gb, price) VALUES
        ('Corsair', 'Vengeance RGB', 'DDR5', 6000, 32, 120.00),
        ('G.Skill', 'Trident Z5', 'DDR5', 6400, 32, 140.00),
        ('Kingston', 'Fury Beast', 'DDR4', 3200, 16, 60.00),
        ('Team Group', 'Delta RGB', 'DDR5', 6000, 16, 80.00);
    `)
	if err != nil {
		return err
	}

	log.Println("✅ Тестовые данные добавлены")
	return nil
}

// Close закрывает подключение к базе данных
func Close() {
	if DB != nil {
		DB.Close()
	}
}
