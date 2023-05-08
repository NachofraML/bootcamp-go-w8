-- Práctica Grupal

-- Resolver las siguientes consignas
-- Tomando la base de datos movies_db.sql, se solicita:
-- Agregar una película a la tabla movies.
INSERT INTO movies VALUES
(null, CURRENT_TIMESTAMP(), null, 'Kimi no Nawa', 10, 1000, '2023-10-10', 240, null);
SELECT * FROM movies_db.movies;
-- Agregar un género a la tabla genres.
INSERT INTO genres VALUES 
(null, CURRENT_TIMESTAMP(), null, 'Anime', 13, 1);
SELECT * FROM movies_db.genres;
-- Asociar a la película del punto 1. genre el género creado en el punto 2.
UPDATE movies SET genre_id = (SELECT id FROM genres WHERE name = 'Anime')
WHERE title = 'Kimi no Nawa';
SELECT * FROM movies_db.movies;
-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
UPDATE actors SET favorite_movie_id = (SELECT id FROM movies WHERE title = 'Kimi no Nawa')
WHERE first_name = 'Leonardo' AND last_name = 'Di Caprio';
SELECT * FROM movies_db.actors;
-- Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE temp_movies_copy SELECT * FROM movies_db.movies;
SELECT * FROM movies_db.temp_movies_copy;
-- Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM temp_movies_copy WHERE awards < 5;
SELECT * FROM movies_db.temp_movies_copy;
-- Obtener la lista de todos los géneros que tengan al menos una película.
SELECT DISTINCT g.name FROM genres g INNER JOIN movies m ON m.genre_id = g.id;
-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT a.*, m.awards FROM actors a INNER JOIN movies m ON m.id = a.favorite_movie_id
WHERE m.awards > 3;
-- Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX movie_name ON movies(title);
DROP INDEX movie_name ON movies; 
-- OR
ALTER TABLE movies_db.movies
ADD INDEX title_movie_index (title ASC) VISIBLE;
-- Chequee que el índice fue creado correctamente.
SHOW INDEXES FROM movies;
-- En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.

-- Hay una mejora notable en consultas que utilizen de filtro el titulo de una pelicula, ya que ahora con el indice no es necesario hacer un full table scan y no creo que moleste al CRUD de la tabla ya que la insercion, 
-- actualizacion y eliminacion de peliculas no es algo lo suficientemente recurrente como para que sea mas necesaria su rapidez de alteracion contrastada contra su rapides de scaneo.
UPDATE movies SET genre_id = (SELECT id FROM genres WHERE name = 'Anime')
WHERE title = 'Kimi no Nawa';
SELECT * FROM movies_db.movies WHERE title = 'Kimi no Nawa';
-- ¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
-- Quizas crearia indices en todos los genre_id de las tablas por si es necesario una mayor rapidez al filtrar generos, 
-- ya que normalmente en todos los servicios de streaming los generos son los filtros de busqueda mas utilizados por los usuarios.

