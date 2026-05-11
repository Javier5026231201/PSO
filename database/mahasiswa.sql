CREATE DATABASE mahasiswa_app;
USE mahasiswa_app;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE notes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE schedules (
    id INT AUTO_INCREMENT PRIMARY KEY,
    course VARCHAR(255),
    day VARCHAR(50),
    time VARCHAR(50),
    room VARCHAR(100)
);

INSERT INTO users(username, password)
VALUES('admin', '12345');