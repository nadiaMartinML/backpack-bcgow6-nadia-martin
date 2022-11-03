/* Ejercicio 4
Plantear 10 consultas SQL que se podrían realizar a la base de datos. Expresar las sentencias */
/*Mostrar los nombres de los clientes de caba */
SELECT nombre FROM cliente WHERE provincia = "caba";
/*Mostrar los nombres y sus cumpleaños de los clientes que nacieron dentro del rango 1984 y 2002 */
SELECT nombre, fecha_nacimiento FROM cliente WHERE fecha_nacimiento BETWEEN "1984-01-01" AND "2002-12-31";
/*Mostrar los clientes que sean mayores de 35 años y vivan en caba */
SELECT * FROM cliente WHERE year(fecha_nacimiento) < 1987 AND provincia = "caba";
/*Mostrar los nombres de los clientes ordenados por dni en forma descendente */
SELECT nombre FROM cliente ORDER BY dni DESC;
/*Mostrar los clientes que tengan el plan menor a 4 */
SELECT * FROM cliente WHERE id_pack < 4;
/*Mostrar los planes que tengan un descuento mayor a 40 */
SELECT * FROM pack WHERE descuento > 40;
/*Mostrar los planes que tengan un descuento mayor a 30 y que salga menos que 5000 */
SELECT id_pack FROM pack WHERE descuento > 30 AND precio < 5000;
/*Mostrar los planes ordenados por velocidad en forma ascendete */
SELECT * FROM pack ORDER BY velocidad ASC;
/*Mostrar los nombres de los clientes que no sean de caba */
SELECT nombre FROM cliente WHERE provincia NOT LIKE "caba"; 
/*Mostrar los nombres de los primeros 3 clientes que vivan en caba */
SELECT nombre FROM cliente WHERE ciudad = "caba" LIMIT 3;