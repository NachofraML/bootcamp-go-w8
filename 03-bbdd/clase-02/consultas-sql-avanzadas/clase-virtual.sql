-- ¿A qué se denomina JOIN en una base de datos y para qué se utiliza?
-- Se utiliza para obtener datos de varias tablas relacionadas entre sí. Consiste en combinar datos de una tabla con datos de la otra tabla, a partir de una o varias condiciones en común.
-- Explicar dos tipos de JOIN.
-- Inner Join se utiliza para traer los datos relacionados de dos o más tablas.
-- Left Join se utiliza para traer los datos de la tabla izquierda más los relacionados de la tabla derecha.
-- ¿Para qué se utiliza el GROUP BY?
-- Agrupa los resultados según las columnas indicadas. 
-- Genera un solo registro por cada grupo de filas que compartan las columnas indicadas.
-- Reduce la cantidad de filas de la consulta.
-- Se suele utilizar en conjunto con funciones de agregación, para obtener datos resumidos y agrupados por las columnas que se necesiten.
-- ¿Para qué se utiliza el HAVING? 
-- La cláusula HAVING se utiliza para incluir condiciones con algunas funciones SQL.
-- Afecta a los resultados traidos por Group By.

-- Escribir una consulta genérica para cada uno de los siguientes diagramas:

-- SELECT * FROM tabla as t INNER JOIN tabla_2 as t2 ON t2.id = t.id;
-- SELECT * FROM tabla as t LEFT JOIN tabla_2 as t2 ON t2.id = t.id;


-- Segunda Parte
-- Se propone realizar las siguientes consultas a la base de datos movies_db.sql trabajada en la primera clase.
-- Importar el archivo movies_db.sql desde PHPMyAdmin o MySQL Workbench y resolver las siguientes consultas:
-- Mostrar el título y el nombre del género de todas las series.
SELECT title, g.name FROM series s INNER JOIN genres g ON g.id = s.genre_id;
-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT e.title, a.first_name, a.last_name FROM episodes e 
INNER JOIN actor_episode ae ON ae.episode_id = e.id
INNER JOIN actors a ON a.id = ae.actor_id;
-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT s.title, COUNT(sea.id) `Cantidad de temporadas` FROM series s INNER JOIN seasons sea ON sea.serie_id = s.id
GROUP BY s.title;
-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT g.name, COUNT(m.id) `Cantidad de peliculas` FROM genres g INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.name
HAVING `Cantidad de peliculas` >= 3;
-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT a.first_name, a.last_name FROM actors a
INNER JOIN actor_movie am ON am.actor_id = a.id
INNER JOIN movies m ON m.id = am.movie_id
WHERE m.title LIKE 'La Guerra de las galaxias%'
