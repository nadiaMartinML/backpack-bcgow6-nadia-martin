/*Ejercitacion clase 2A asincronica */
/* Mostrar el título y el nombre del género de todas las series.*/
SELECT s.title, g.name FROM series s
INNER JOIN genres g
ON g.id = s.genre_id;

/* Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.*/
SELECT e.title, CONCAT(a.first_name, " ", a.last_name) AS "Nombre y Apellido" FROM episodes e
INNER JOIN actor_episode ae ON ae.episode_id = e.id
INNER JOIN actors a ON a.id = ae.actor_id;

/* Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.*/
SELECT s.title, COUNT(t.number) AS "Total Temporadas" FROM series s
INNER JOIN seasons t ON t.serie_id = s.id
GROUP BY s.title;

/* Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, 
siempre que sea mayor o igual a 3.*/
SELECT g.name, COUNT(m.title) AS "Total Peliculas" FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.name
HAVING COUNT(m.title) >= 3;

/* Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas 
de la guerra de las galaxias y que estos no se repitan.*/
SELECT DISTINCT CONCAT(a.first_name, " ", a.last_name) AS "Nombre y Apellido" FROM actors a
INNER JOIN actor_movie am ON am.actor_id = a.id
INNER JOIN movies m ON m.id = am.movie_id
WHERE m.title LIKE "%La Guerra de las galaxias%";