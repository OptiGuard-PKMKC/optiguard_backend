CREATE TABLE IF NOT EXISTS user_roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(20) NOT NULL UNIQUE
);

INSERT INTO user_roles (role_name) VALUES
('guest'),
('admin'),
('patient'),
('doctor');

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
    CONSTRAINT fk_users_role_id FOREIGN KEY (role_id) REFERENCES user_roles(id)
);

CREATE TABLE IF NOT EXISTS funduses (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL,
    image_path TEXT NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    status VARCHAR(255),
    condition VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_funduses_patient_id FOREIGN KEY (patient_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS fundus_details (
    id SERIAL PRIMARY KEY,
    fundus_id INTEGER NOT NULL,
    disease VARCHAR(255) NOT NULL,
    confidence_score FLOAT NOT NULL,
    description TEXT DEFAULT '',
    CONSTRAINT fk_fundus_details_fundus_id FOREIGN KEY (fundus_id) REFERENCES funduses(id)
);

CREATE TABLE IF NOT EXISTS fundus_feedbacks (
    id SERIAL PRIMARY KEY,
    fundus_id INTEGER NOT NULL,
    doctor_id INTEGER NOT NULL,
    notes TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_fundus_feedbacks_fundus_id FOREIGN KEY (fundus_id) REFERENCES funduses(id),
    CONSTRAINT fk_fundus_feedbacks_doctor_id FOREIGN KEY (doctor_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS doctor_profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    specialization VARCHAR(255) NOT NULL,
    str_number VARCHAR(255) NOT NULL,
    bio_desc TEXT DEFAULT '',
    work_years INTEGER DEFAULT 0,
    rating INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_doctor_profiles_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS doctor_schedules (
    id SERIAL PRIMARY KEY,
    doctor_id INTEGER NOT NULL,
    day_of_week INTEGER NOT NULL,
    start_hour TIME NOT NULL,
    end_hour TIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_doctor_schedules_doctor_id FOREIGN KEY (doctor_id) REFERENCES doctor_profiles(id)
);

CREATE TABLE IF NOT EXISTS doctor_practices (
    id SERIAL PRIMARY KEY,
    doctor_id INTEGER NOT NULL,
    city VARCHAR(255) NOT NULL,
    province VARCHAR(255) NOT NULL,
    office_name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_doctor_practices_doctor_id FOREIGN KEY (doctor_id) REFERENCES doctor_profiles(id)
);

CREATE TABLE IF NOT EXISTS doctor_educations (
    id SERIAL PRIMARY KEY,
    doctor_id INTEGER NOT NULL,
    degree VARCHAR(255) NOT NULL,
    school_name VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_doctor_educations_doctor_id FOREIGN KEY (doctor_id) REFERENCES doctor_profiles(id)
);

CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL,
    doctor_id INTEGER NOT NULL,
    date DATE NOT NULL,
    start_hour TIME NOT NULL,
    end_hour TIME NOT NULL,
    status VARCHAR(255) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_appointments_patient_id FOREIGN KEY (patient_id) REFERENCES users(id),
    CONSTRAINT fk_appointments_doctor_id FOREIGN KEY (doctor_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS health_facilities (
    id SERIAL PRIMARY KEY,
    facility_name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    province VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    adaptor_quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE facility_schedules (
	id SERIAL PRIMARY KEY,
	facility_id INTEGER NOT NULL,
	day_of_week INTEGER NOT NULL,
	start_hour TIME NOT NULL,
	end_hour TIME NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_facility_schedules_facility_id FOREIGN KEY (facility_id) REFERENCES health_facilities(id)
);

CREATE TABLE IF NOT EXISTS adaptors (
    id SERIAL PRIMARY KEY,
    facility_id INTEGER NOT NULL,
    device_code VARCHAR(255) NOT NULL,
    used BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_adaptors_facility_id FOREIGN KEY (facility_id) REFERENCES health_facilities(id)
);

CREATE TABLE IF NOT EXISTS user_adaptors (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    facility_id INTEGER NOT NULL,
    adaptor_id INTEGER NOT NULL,
    schedule_id INTEGER NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_adaptors_user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_user_adaptors_adaptor_id FOREIGN KEY (adaptor_id) REFERENCES adaptors(id),
    CONSTRAINT fk_user_adaptors_schedule_id FOREIGN KEY (schedule_id) REFERENCES facility_schedules(id),
    CONSTRAINT fk_user_adaptors_facility_id FOREIGN KEY (facility_id) REFERENCES health_facilities(id)
);

CREATE TABLE IF NOT EXISTS chat_rooms (
    id SERIAL PRIMARY KEY,
    doctor_id INTEGER NOT NULL,
    patient_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_chat_rooms_doctor_id FOREIGN KEY (doctor_id) REFERENCES users(id),
    CONSTRAINT fk_chat_rooms_patient_id FOREIGN KEY (patient_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS chat_messages (
    id SERIAL PRIMARY KEY,
    chat_room_id INTEGER NOT NULL,
    sender_id INTEGER NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_chat_messages_chat_room_id FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(id),
    CONSTRAINT fk_chat_messages_sender_id FOREIGN KEY (sender_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_notifications_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);
