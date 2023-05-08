-- Práctica Individual

-- Escenario 
-- Luego de un análisis realizado en un sistema de facturación, se ha detectado un mal diseño en la base de datos. La misma, cuenta con una tabla facturas que almacena datos de diferente naturaleza. 
-- Como se puede observar, la tabla cuenta con datos que podrían ser normalizados y separados en diferentes entidades.

-- Ejercicio 
-- Se solicita para el escenario anterior: 
-- Aplicar reglas de normalización y elaborar un modelo de DER que alcance la tercera forma normal (3FN).
-- Describir con sus palabras cada paso de la descomposición y aplicación de las reglas para visualizar el planteo realizado.

-- 1FN
-- Elimino repeticion de datos.
-- A la tabla actual le cambiaria el nombre a Factura y dejaria solo estos campos: id_factura, fecha_factura, cantidad, importe.

-- 2FN
-- Separo dependencia parciales
-- Separaria los campos cliente y articulo en una tabla Cliente y otra Articulo, asi no es necesario en cada columna identificar a que entidad pertenece.
-- Los campos de cliente serian id_cliente, nombre, apellido, direccion y los de articulo serian articulo_id, nombre, descripcion, precio (y aca podria ir IVA si no se crea una tabla a parte,
-- en caso de que se cree una tabla a parte aca deberia ir la id de categoria del producto).

-- 3FN
-- Separo subconjuntos de datos que podrian ser mejor distribuidos
-- Lo que haria es separar la forma de pago en una tabla la cual contenga cada forma de pago y su id correspondiente.
-- Tambien se podria separar el IVA y conectarlo con una tabla que sea por ejemplo de categorias de producto, para no tener que repetir el IVA en cada producto (la tabla de IVA quedaria con la id de categoria de producto a la que pertenece).

