
/*EJECUTAR EN EL WORKBENCH DE MYSQL*/
CREATE DATABASE proyecto_desarrollo_sw;

USE proyecto_desarrollo_sw;

CREATE TABLE usuarios (
id INT AUTO_INCREMENT PRIMARY KEY,
nombre VARCHAR(100) NOT NULL,
email VARCHAR(100) NOT NULL UNIQUE,
password_hash VARCHAR(255) NOT NULL,
rol VARCHAR(20) NOT NULL
);

CREATE TABLE eventos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descripcion TEXT,
    fecha DATE NOT NULL,
    horario TIME NOT NULL,
    duracion INT NOT NULL,
    ubicacion VARCHAR(255) NOT NULL,
    capacidad INT NOT NULL,
    precio DECIMAL(10,2) NOT NULL,
    categoria VARCHAR(100),
    imagen_url VARCHAR(255),
    estado VARCHAR(20) NOT NULL
);