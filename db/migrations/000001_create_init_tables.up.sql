CREATE TABLE IF NOT EXISTS user_roles (
    id int PRIMARY KEY,
    role_name VARCHAR(20) NOT NULL UNIQUE
);

INSERT INTO user_roles (id, role_name) VALUES 
(1, 'guest'),
(2, 'admin'),
(3, 'patient'),
(4, 'doctor');

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(24) DEFAULT '' UNIQUE,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    birthdate DATE DEFAULT CURRENT_DATE,
    gender VARCHAR(6) DEFAULT '',
    city VARCHAR(255) DEFAULT '',
    province VARCHAR(255) DEFAULT '',
    address TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES user_roles(id)
);

CREATE TABLE IF NOT EXISTS funduses (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL,
    image_url TEXT NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    status VARCHAR(255),
    condition VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_patient_id FOREIGN KEY (patient_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS fundus_details (
    id SERIAL PRIMARY KEY,
    fundus_id INTEGER NOT NULL,
    disease VARCHAR(255) NOT NULL,
    confidence_score FLOAT NOT NULL,
    description TEXT DEFAULT '',
    CONSTRAINT fk_fundus_id FOREIGN KEY (fundus_id) REFERENCES funduses(id)
);

CREATE TABLE IF NOT EXISTS fundus_feedbacks (
    id SERIAL PRIMARY KEY,
    fundus_id INTEGER NOT NULL,
    doctor_id INTEGER NOT NULL,
    notes TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_fundus_id FOREIGN KEY (fundus_id) REFERENCES funduses(id),
    CONSTRAINT fk_doctor_id FOREIGN KEY (doctor_id) REFERENCES users(id)
);