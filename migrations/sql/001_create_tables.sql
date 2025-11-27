-- Производители (один ко многим)
        CREATE TABLE vendors (
            id SERIAL PRIMARY KEY,
            name VARCHAR(50) UNIQUE NOT NULL
        );
        INSERT INTO vendor (name) VALUES
        ('Intel'),
        ('AMD'),
        ('NVIDIA'),
        ('GIGABYTE'),
        ('MSI'),
        ('ASUS'),
        ('KINGSTONE'),
        ('SAMSUNG'),
        ('HP'),
        ('DEEPCOOL'),
        ('EXEGATE');

-- Категории компонентов 
        CREATE TABLE categories (
            id SERIAL PRIMARY KEY,
            name VARCHAR(50) UNIQUE NOT NULL,
            description TEXT
        );
        INSERT INTO categories (name, description) VALUES
        ('Gaming','Maximum performance for a comfortable gaming experience'),
        ('Office','Budget version for everyday PC use'),
        ('Workstation','Components for server machine assembly');


-- Сокеты процессоров
        CREATE TABLE sockets (
            id SERIAL PRIMARY KEY,
            name VARCHAR(20) UNIQUE NOT NULL, -- AM5, LGA1700
        );
        INSERT INTO sockets (name) VALUES
        ('AM4'),
        ('AM5'),
        ('LGA1851'),
        ('LGA1700'),
        ('LGA1200');

-- Форм-факторы
        CREATE TABLE form_factors (
            id SERIAL PRIMARY KEY,
            name VARCHAR(20) UNIQUE NOT NULL
        );
        INSERT INTO form_factors (name) VALUES
        ('ATX'),
        ('mATX'),
        ('ITX');

-- Типы памяти
        CREATE TABLE memory_types (
            id SERIAL PRIMARY KEY,
            name VARCHAR(10) UNIQUE NOT NULL  -- DDR4, DDR5, GDDR6
        );
        INSERT INTO memory_types (name) VALUES
        ('DDR4'),
        ('DDR5'),
        ('GDDR5'),
        ('GDDR6X'),
        ('GDDR7');

-- Процессоры
        CREATE TABLE cpus (
            id SERIAL PRIMARY KEY,
            vendor_id INTEGER REFERENCES vendors(id),
            model VARCHAR(100) NOT NULL,
            socket_id INTEGER REFERENCES sockets(id),
            core_count INTEGER NOT NULL,
            thread_count INTEGER NOT NULL,
            base_clock DECIMAL(4,2) NOT NULL, -- GHz
            boost_clock DECIMAL(4,2),          -- GHz
            l3_cache INTEGER,                  -- MB
            tdp INTEGER NOT NULL,              -- W
            launch_date INTEGER,
            price DECIMAL(10,2) NOT NULL,
            UNIQUE(vendor_id, model)
        );
        INSERT INTO cpus (vendor_id, model, socket_id, core_count, thread_count, base_clock, boost_clock, l3_cache, tpd, launch_date, price) VALUES
        ('2', 'Ryzen 7 7800x3d', '2', '8', '16', '4.2', '5.0', '96', '120', '2023', '1150.00');

-- Видеокарты
        CREATE TABLE gpus (
            id SERIAL PRIMARY KEY,
            vendor_id INTEGER REFERENCES vendors(id),
            model VARCHAR(100) NOT NULL,
            vram_size INTEGER NOT NULL,        -- GB
            memory_type_id INTEGER REFERENCES memory_types(id),
            base_clock INTEGER,                -- MHz
            boost_clock INTEGER,               -- MHz
            tdp INTEGER NOT NULL,              -- W
            length INTEGER,                    -- mm
            power_connectors VARCHAR(50),      -- 8-pin, 6+8-pin
            price DECIMAL(10,2) NOT NULL,
            UNIQUE(vendor_id, model)
        );
        INSERT INTO gpus (vendor_id, model, vram_size, memory_type_id, base_clock, boost_clock, tpd, length, power_connectors, price) VALUES
        ('4', 'GeForce RTX 3070 Ti', '8', '4', '1580', '1830', '290', '320', '8+8-pin', '2871.00');

-- Процессорные куллеры
        CREATE TABLE cpu_coolers (
            id SERIAL PRIMARY KEY,
            vendor_id INTEGER REFERENCES vendors(id),
            model VARCHAR(100) NOT NULL,
            cooler_type VARCHAR(20) NOT NULL,  -- air, liquid
            radiator_size INTEGER,             -- mm (для СЖО)
            height INTEGER,                    -- mm
            supported_sockets JSONB,           -- массив сокетов ["AM5", "LGA1700"]
            tdp INTEGER,                       -- W (рассеиваемая мощность)
            noise_level DECIMAL(4,2),          -- dB
            price DECIMAL(10,2) NOT NULL,
            UNIQUE(vendor_id, model)
        );

-- Корпусные куллеры
        CREATE TABLE case_fans (
            id SERIAL PRIMARY KEY,
            vendor_id INTEGER REFERENCES vendors(id),
            model VARCHAR(100) NOT NULL,
            size INTEGER NOT NULL,             -- mm (120, 140)
            airflow DECIMAL(5,2),              -- CFM
            static_pressure DECIMAL(4,2),      -- mm H2O
            noise_level DECIMAL(4,2),          -- dB
            pwm BOOLEAN DEFAULT TRUE,          -- PWM управление
            rgb BOOLEAN DEFAULT FALSE,
            price DECIMAL(10,2) NOT NULL,
            UNIQUE(vendor_id, model)
        );