/* 
Ejercicio 2 
Una vez modelada y planteada la base de datos, responder a las siguientes preguntas:

a. ¿Cuál es la primary key para la tabla de clientes? Justificar respuesta
	La PK es el id, ya que no es seguro utilizar el campo DNI como una primary key ya que existen DNI duplicados.
b. ¿Cuál es la primary key para la tabla de planes de internet? Justificar respuesta.
	La PK es la identificación del plan, ya que es algo que sale de la compañía y son ellos los creadores de ese 
    campo, por ende esta en nuestra alcance saber que no es un campo que podría estar duplicado.
c. ¿Cómo serían las relaciones entre tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de 
qué tabla hace referencia dicha foreign key? Justificar respuesta.
	La relación entre tablas sería que un cliente tiene un plan, pero un plan puede tener varios clientes.
	La FK debería estar en la tabla de clientes, indicando el identificador del plan.
*/