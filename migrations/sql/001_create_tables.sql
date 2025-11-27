-- 1. Производители (один ко многим)
CREATE TABLE IF NOT EXISTS vendors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);
-- INSERT INTO vendors (name) VALUES
-- ('Intel'),
-- ('AMD'),
-- ('NVIDIA'),
-- ('GIGABYTE'),
-- ('MSI'),
-- ('ASUS'),
-- ('KINGSTONE'),
-- ('SAMSUNG'),
-- ('HP'),
-- ('DEEPCOOL'),
-- ('EXEGATE'),
-- ('KINGBANK');

-- 2. Категории компонентов 
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT
);
-- INSERT INTO categories (name, description) VALUES
-- ('Gaming','Maximum performance for a comfortable gaming experience'),
-- ('Office','Budget version for everyday PC use'),
-- ('Workstation','Components for server machine assembly');

-- 3. Сокеты процессоров
CREATE TABLE IF NOT EXISTS sockets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) UNIQUE NOT NULL
);
-- INSERT INTO sockets (name) VALUES
-- ('AM4'),
-- ('AM5'),
-- ('LGA1851'),
-- ('LGA1700'),
-- ('LGA1200');

-- 4. Форм-факторы
CREATE TABLE IF NOT EXISTS form_factors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) UNIQUE NOT NULL
);
-- INSERT INTO form_factors (name) VALUES
-- ('ATX'),
-- ('mATX'),
-- ('ITX');

-- 5. Типы памяти
CREATE TABLE IF NOT EXISTS memory_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) UNIQUE NOT NULL
);
-- INSERT INTO memory_types (name) VALUES
-- ('DDR4'),
-- ('DDR5'),
-- ('GDDR5'),
-- ('GDDR6X'),
-- ('GDDR7');

-- 6. Процессоры
CREATE TABLE IF NOT EXISTS cpus (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER REFERENCES vendors(id),
    model VARCHAR(100) NOT NULL,
    socket_id INTEGER REFERENCES sockets(id),
    core_count INTEGER NOT NULL,
    thread_count INTEGER NOT NULL,
    base_clock DECIMAL(4,2) NOT NULL,
    boost_clock DECIMAL(4,2),
    l3_cache INTEGER,
    tdp INTEGER NOT NULL,
    launch_date INTEGER,
    price DECIMAL(10,2) NOT NULL,
    UNIQUE(vendor_id, model)
);
-- INSERT INTO cpus (vendor_id, model, socket_id, core_count, thread_count, base_clock, boost_clock, l3_cache, tdp, launch_date, price) VALUES
-- (2, 'Ryzen 7 7800x3d', 2, 8, 16, 4.2, 5.0, 96, 120, 2023, 1150.00);

-- 7. Видеокарты
CREATE TABLE IF NOT EXISTS gpus (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER REFERENCES vendors(id),
    model VARCHAR(100) NOT NULL,
    vram_size INTEGER NOT NULL,
    memory_type_id INTEGER REFERENCES memory_types(id),
    base_clock INTEGER,
    boost_clock INTEGER,
    tdp INTEGER NOT NULL,
    length INTEGER,
    power_connectors VARCHAR(50),
    price DECIMAL(10,2) NOT NULL,
    UNIQUE(vendor_id, model)
);
-- INSERT INTO gpus (vendor_id, model, vram_size, memory_type_id, base_clock, boost_clock, tdp, length, power_connectors, price) VALUES
-- (3, 'GeForce RTX 3070 Ti', 8, 4, 1580, 1830, 290, 320, '8+8-pin', 2871.00);

-- 8. Материнские платы
CREATE TABLE IF NOT EXISTS motherboards (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER REFERENCES vendors(id),
    model VARCHAR(100) NOT NULL,
    socket_id INTEGER REFERENCES sockets(id),
    chipset VARCHAR(50) NOT NULL,
    form_factor_id INTEGER REFERENCES form_factors(id),
    memory_slots INTEGER NOT NULL,
    max_memory_gb INTEGER,
    max_memory_speed INTEGER,
    memory_type_id INTEGER REFERENCES memory_types(id),
    price DECIMAL(10,2) NOT NULL,
    UNIQUE(vendor_id, model)
);
-- INSERT INTO motherboards (vendor_id, model, socket_id, chipset, form_factor_id, memory_slots, max_memory_gb, max_memory_speed, memory_type_id, price) VALUES
-- (6, 'TUF Gaming X870-Plus WiFi', 2, 'X870', 1, 4, 192, 8000, 2, 1260.10);

-- 9. Оперативная память
CREATE TABLE IF NOT EXISTS rams (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER REFERENCES vendors(id),
    model VARCHAR(100) NOT NULL,
    memory_type_id INTEGER REFERENCES memory_types(id),
    speed_mhz INTEGER NOT NULL,
    capacity_gb INTEGER NOT NULL,
    latency VARCHAR(20),
    modules_count INTEGER NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    UNIQUE(vendor_id, model)
);
-- INSERT INTO rams (vendor_id, model, memory_type_id, speed_mhz, capacity_gb, latency, modules_count, price) VALUES
-- (12, 'KRRB 2x16ГБ K5.01.FLM5ED9402', 2, 6000, 32, '30-36-36-76', 2, 1125.00);

-- 10. Накопители
CREATE TABLE IF NOT EXISTS storages (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER REFERENCES vendors(id),
    model VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL,
    form_factor VARCHAR(20),
    capacity_gb INTEGER NOT NULL,
    read_speed INTEGER,
    write_speed INTEGER,
    price DECIMAL(10,2) NOT NULL,
    UNIQUE(vendor_id, model)
);
-- INSERT INTO storages (vendor_id, model, type, form_factor, capacity_gb, read_speed, write_speed, price) VALUES
-- (8, 'SSD 980 PRO 1TB', 'SSD', 'M.2', 1000, 7000, 5000, 245.50);

-- 11. Блоки питания
CREATE TABLE IF NOT EXISTS psus (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER REFERENCES vendors(id),
    model VARCHAR(100) NOT NULL,
    wattage INTEGER NOT NULL,
    form_factor_id INTEGER REFERENCES form_factors(id),
    efficiency_rating VARCHAR(20),
    modularity VARCHAR(20),
    price DECIMAL(10,2) NOT NULL,
    UNIQUE(vendor_id, model)
);
-- INSERT INTO psus (vendor_id, model, wattage, form_factor_id, efficiency_rating, modularity, price) VALUES
-- (4, 'GP-P850GM', 850, 1, '80+ Gold', 'Semi', 320.00);

-- 12. Корпуса
CREATE TABLE IF NOT EXISTS cases (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER REFERENCES vendors(id),
    model VARCHAR(100) NOT NULL,
    form_factor_id INTEGER REFERENCES form_factors(id),
    max_gpu_length INTEGER,
    max_cpu_cooler_height INTEGER,
    psu_support BOOLEAN DEFAULT TRUE,
    price DECIMAL(10,2) NOT NULL,
    UNIQUE(vendor_id, model)
);
-- INSERT INTO cases (vendor_id, model, form_factor_id, max_gpu_length, max_cpu_cooler_height, psu_support, price) VALUES
-- (10, 'MATREXX 55 MESH', 1, 350, 165, true, 85.00);

-- 13. Кулеры процессоров
CREATE TABLE IF NOT EXISTS cpu_coolers (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER REFERENCES vendors(id),
    model VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL,
    radiator_size INTEGER,
    height INTEGER,
    noise_level DECIMAL(4,2),
    tdp INTEGER,
    price DECIMAL(10,2) NOT NULL,
    UNIQUE(vendor_id, model)
);
-- INSERT INTO cpu_coolers (vendor_id, model, type, radiator_size, height, noise_level, tdp, price) VALUES
-- (10, 'AK620', 'air', NULL, 160, 28.5, 260, 75.00);

-- 14. Пользователи
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- INSERT INTO users (username, email, password_hash) VALUES
-- ('test_user', 'user@test.com', 'hashed_password_123');

-- 15. Сборки ПК
CREATE TABLE IF NOT EXISTS builds (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    name VARCHAR(100) NOT NULL,
    total_price DECIMAL(10,2) DEFAULT 0,
    total_tdp INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- INSERT INTO builds (user_id, name, total_price, total_tdp) VALUES
-- (1, 'My Gaming PC', 0, 0);