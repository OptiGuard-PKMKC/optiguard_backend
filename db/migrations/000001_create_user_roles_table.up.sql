CREATE TABLE IF NOT EXISTS user_roles (
    id int PRIMARY KEY,
    role_name VARCHAR(20) NOT NULL UNIQUE
);

INSERT INTO user_roles (id, role_name) VALUES 
(1, 'guest'),
(2, 'admin'),
(3, 'patient'),
(4, 'doctor');