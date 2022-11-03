/*Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.*/
SELECT e.nombre, e.puesto, d.localidad FROM empleado e INNER JOIN departamento d 
ON d.dpto_nro = e.depto_nro;
/*Visualizar los departamentos con más de cinco empleados.*/
SELECT d.nombre_depto FROM departamento d INNER JOIN empleado e 
ON e.depto_nro = d.dpto_nro GROUP BY d.nombre_depto HAVING COUNT(e.nombre) > 5;
/*Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.*/
SELECT e.nombre, e.salario, d.nombre_depto FROM empleado e
INNER JOIN departamento d ON d.dpto_nro = e.depto_nro
WHERE d.dpto_nro = (SELECT depto_nro FROM empleado WHERE CONCAT(nombre, " ", apellido) = "Mito Barchuk");
/*Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.*/
SELECT e.* FROM empleado e INNER JOIN departamento d ON d.dpto_nro = e.depto_nro 
WHERE d.nombre_depto = "contabilidad" ORDER BY e.nombre;
/*Mostrar el nombre del empleado que tiene el salario más bajo.*/
SELECT e.nombre FROM empleado e ORDER BY salario LIMIT 1;
/*Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.*/
SELECT e.* FROM empleado e INNER JOIN departamento d ON d.dpto_nro = e.depto_nro
WHERE nombre_depto = "ventas" ORDER BY salario DESC LIMIT 1;