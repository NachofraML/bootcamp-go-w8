-- Ejercicio 1
-- Se tiene el siguiente DER que corresponde al esquema que presenta la base de datos de una “biblioteca”.



-- En base al mismo, plantear las consultas SQL para resolver los siguientes requerimientos:
-- Listar los datos de los autores.
SELECT * FROM autor;
-- Listar nombre y edad de los estudiantes
SELECT nombre, edad FROM estudiante;
-- ¿Qué estudiantes pertenecen a la carrera informática?
SELECT * FROM estudiante
WHERE carrera = 'informatica';
-- ¿Qué autores son de nacionalidad francesa o italiana?
SELECT * FROM autor
WHERE nacionalidad IN ('francesa', 'italiana');
-- ¿Qué libros no son del área de internet?
SELECT * FROM libro
WHERE area != 'internet';
-- Listar los libros de la editorial Salamandra.
SELECT * FROM libro
WHERE editorial = 'Salamandra';
-- Listar los datos de los estudiantes cuya edad es mayor al promedio.
SELECT * FROM estudiante
WHERE edad > (SELECT AVG(edad) FROM estudiante);
-- Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
SELECT nombre FROM estudiante
WHERE apellido LIKE 'G%';
-- Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
SELECT a.nombre FROM autor a 
INNER JOIN libroautor la ON la.idAutor = a.idAutor
INNER JOIN libro l ON l.idLibro= la.idLibro
WHERE l.titulo = 'El Universo: Guía de viaje';
-- ¿Qué libros se prestaron al lector “Filippo Galli”?
SELECT l.titulo FROM prestamo p
INNER JOIN estudiante e ON e.idLector = p.idLector
INNER JOIN libro l ON l.idLibro = p.idLibro
WHERE e.nombre = 'Filippo' AND e.apellido = 'Galli';
-- Listar el nombre del estudiante de menor edad.
SELECT nombre FROM estudiante
WHERE edad = (SELECT MIN(edad) FROM estudiante);
-- Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
SELECT DISTINCT e.nombre FROM estudiante e
INNER JOIN prestamo P ON P.idLector = e.idLector
INNER JOIN libro l ON l.idLibro = p.idLibro
WHERE l.area = 'Bases de Datos';
-- Listar los libros que pertenecen a la autora J.K. Rowling.
SELECT l.* FROM autor a 
INNER JOIN libroautor la ON la.idAutor = a.idAutor
INNER JOIN libro l ON l.idLibro= la.idLibro
WHERE a.nombre = 'J.K. Rowling';
-- Listar títulos de los libros que debían devolverse el 16/07/2021.
SELECT l.titulo FROM prestamo p
INNER JOIN libro l ON l.idLibro = p.idLibro
WHERE p.fechadevolucion = '2021-07-16';


-- Ejercicio 2
-- Implementar la base de datos en PHPMyAdmin o MySQL Workbench, cargar cinco registros en cada tabla y probar algunas consultas planteadas en el Ejercicio 1. 

CREATE DATABASE biblioteca;

USE biblioteca;

CREATE TABLE Estudiante (
  idLector INT PRIMARY KEY,
  Nombre VARCHAR(50),
  Apellido VARCHAR(50),
  Direccion VARCHAR(100),
  Carrera VARCHAR(50),
  Edad INT
);

CREATE TABLE Libro (
  idLibro INT PRIMARY KEY,
  Titulo VARCHAR(100),
  Editorial VARCHAR(50),
  Area VARCHAR(50)
);

CREATE TABLE Prestamo (
  idPrestamo INT PRIMARY KEY AUTO_INCREMENT,
  idLector INT,
  idLibro INT,
  FechaPrestamo DATE,
  FechaDevolucion DATE,
  Devuelto BOOLEAN,
  FOREIGN KEY (idLector) REFERENCES Estudiante(idLector),
  FOREIGN KEY (idLibro) REFERENCES Libro(idLibro)
);

CREATE TABLE Autor (
  idAutor INT PRIMARY KEY,
  Nombre VARCHAR(50),
  Nacionalidad VARCHAR(50)
);

CREATE TABLE LibroAutor (
  idAutor INT,
  idLibro INT,
  FOREIGN KEY (idAutor) REFERENCES Autor(idAutor),
  FOREIGN KEY (idLibro) REFERENCES Libro(idLibro),
  PRIMARY KEY (idAutor, idLibro)
);


INSERT INTO Estudiante (idLector, Nombre, Apellido, Direccion, Carrera, Edad) VALUES
(1, 'Juan', 'Perez', 'Calle 123', 'Informática', 20),
(2, 'María', 'García', 'Avenida 456', 'Ingeniería', 22),
(3, 'Pedro', 'González', 'Calle 789', 'Informática', 19),
(4, 'Ana', 'Rodríguez', 'Avenida 012', 'Arquitectura', 21),
(5, 'Luis', 'Gómez', 'Calle 345', 'Medicina', 23),
(6, 'Filippo', 'Galli', 'Via della Vittoria, 65', 'Ingeniería', 24),
(7, 'Ana', 'González', 'Calle 456', 'Ingeniería', 21),
(8, 'José', 'Martínez', 'Avenida 789', 'Medicina', 20),
(9, 'Marina', 'Silva', 'Calle 012', 'Arquitectura', 22),
(10, 'Pablo', 'López', 'Avenida 345', 'Informática', 19),
(11, 'Carla', 'Gómez', 'Calle 678', 'Derecho', 23);

INSERT INTO Libro (idLibro, Titulo, Editorial, Area) VALUES
(1, 'El Universo: Guía de viaje', 'Salamandra', 'Ciencia'),
(2, 'La Sombra del Viento', 'Planeta', 'Ficción'),
(3, 'Harry Potter y la piedra filosofal', 'Bloomsbury', 'Ficción'),
(4, 'La Odisea', 'Alianza Editorial', 'Clásicos'),
(5, 'Base de Datos', 'McGraw-Hill', 'Bases de Datos'),
(6, 'El Principito', 'Emecé Editores', 'Ficción'),
(7, 'La Historia del Tiempo', 'Bantam Books', 'Ciencia'),
(8, 'El Perfume', 'Seix Barral', 'Ficción'),
(9, 'La Iliada', 'Alianza Editorial', 'Clásicos'),
(10, 'Los Pilares de la Tierra', 'Debolsillo', 'Histórica'),
(11, 'Introducción a la Programación', 'Pearson', 'Bases de Datos'),
(12, 'El Psicoanalista', 'B de Bolsillo', 'Ficción');

INSERT INTO Autor (idAutor, Nombre, Nacionalidad) VALUES
(1, 'Stephen Hawking', 'Británica'),
(2, 'Carlos Ruiz Zafón', 'Española'),
(3, 'J.K. Rowling', 'Británica'),
(4, 'Homero', 'Griega'),
(5, 'Abraham Silberschatz', 'Estadounidense'),
(6, 'Antoine de Saint-Exupéry', 'Francés'),
(7, 'Stephen Hawking', 'Británica');

INSERT INTO LibroAutor (idAutor, idLibro) VALUES
(1, 1),
(1, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7);

INSERT INTO Prestamo (idLector, idLibro, FechaPrestamo, FechaDevolucion, Devuelto) VALUES
(1, 1, '2021-04-01', '2021-04-15', true),
(2, 2, '2021-04-03', '2021-04-17', true),
(3, 3, '2021-04-05', '2021-04-19', true),
(4, 4, '2021-04-07', '2021-04-21', true),
(5, 5, '2021-04-09', '2021-04-23', true),
(6, 6, '2023-04-20', '2023-05-04', false),
(6, 7, '2023-04-21', '2023-05-05', false),
(1, 7, '2021-06-21', '2021-07-16', false),
(8, 7, '2022-03-01', '2022-03-15', true),
(9, 8, '2022-03-02', '2022-03-16', true),
(10, 9, '2022-03-03', '2022-03-17', true),
(11, 10, '2022-03-04', '2022-03-18', true),
(11, 12, '2022-03-05', '2022-03-19', true);
