/* EJERCICIO 1
En base al mismo, plantear las consultas SQL para resolver los siguientes requerimientos:
/*Listar los datos de los autores.*/
	SELECT * FROM autor;
/*Listar nombre y edad de los estudiantes*/
	SELECT nombre, edad FROM estudiante;
/*¿Qué estudiantes pertenecen a la carrera informática?*/
	SELECT nombre FROM estudiante WHERE carrera LIKE "%informatica%";
/*¿Qué autores son de nacionalidad francesa o italiana?*/
	SELECT nombre FROM autor WHERE nacionalidad = "francesa" OR "italiana";
/*¿Qué libros no son del área de internet?*/
	SELECT titulo FROM libro WHERE area NOT LIKE "internet";
/*Listar los libros de la editorial Salamandra.*/
	SELECT titulo FROM libro WHERE editorial = "Salamandra";
/*Listar los datos de los estudiantes cuya edad es mayor al promedio.*/
	SELECT * FROM estudiante WHERE edad > AVG(edad);
/*Listar los nombres de los estudiantes cuyo apellido comience con la letra G.*/
	SELECT nombre FROM estudiante WHERE apellido LIKE "G%";
/*Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).*/
	SELECT a.nombre FROM autor a INNER JOIN libroautor lb ON lb.idAutor = a.idAutor 
    INNER JOIN libro l ON l.idLibro = lb.idLibro WHERE l.titulo LIKE "%El Universo: Guía de viaje%";
/*¿Qué libros se prestaron al lector “Filippo Galli”?*/
	SELECT l.titulo FROM libro l INNER JOIN prestamo p ON p.idLibro = l.idLibro
    INNER JOIN estudiante e ON e.idLector = p.idLector WHERE CONCAT(e.nombre, " ", e.apellido) = "Filippo Galli";
/*Listar el nombre del estudiante de menor edad.*/
	SELECT nombre FROM estudiante WHERE edad < 18;
/*Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.*/
	SELECT CONCAT(e.nombre, " ", e.apellido) AS "Nombre y apellido" FROM estudiante e 
    INNER JOIN prestamo p ON p.idLector = e.idLector
    INNER JOIN libro l ON l.idLibro = p.idLibro WHERE area = "Base de Datos"; 
/*Listar los libros que pertenecen a la autora J.K. Rowling.*/
	SELECT l.titulo FROM libro l INNER JOIN libroautor la ON la.idLibro = l.idLibro
    INNER JOIN autor a ON a.idAutor = la.idAutor WHERE a.nombre = "J.K. Rowling";
/*Listar títulos de los libros que debían devolverse el 16/07/2021. */
	SELECT l.titulo FROM libro l INNER JOIN prestamo p ON p.idLibro = l.idLibro WHERE fechaDevolucion = "2021-07-16";
