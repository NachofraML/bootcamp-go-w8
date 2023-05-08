insert into cliente values (null, 1023584, 'Mauricio', 'Gomez', '2000-01-15', 'Buenos Aires', 'Ciudad Capital');
insert into cliente values (null, 1023524, 'Maury', 'Alcocer', '2003-02-11', 'Cordoba', 'Cordoba');
insert into cliente values (null, 233584, 'Flor', 'Jimenez', '1965-09-25', 'Boyaca', 'Ramiriquí');
insert into cliente values (null, 96503254, 'Tatiana', 'Vasquez', '2001-08-30', 'Cundimarca', 'Tabio');
insert into cliente values (null, 51335497, 'Marcos', 'Gomez', '2002-05-20', 'Buenos Aires', 'Ciudad Capital');
insert into cliente values (null, 1023584, 'Soledad', 'Velandia', '1989-05-21', 'San Luis', 'Ciudad de San Luis');
insert into cliente values (null, 1025684, 'Marcos', 'Galperín', '1960-05-21', 'Buenos Aires', 'Ciudad Capital');
insert into cliente values (null, 152684, 'Lionel', 'Messi', '2000-06-16', 'Rosario', 'Rosario');
insert into cliente values (null, 1023584, 'Luis', 'Diaz', '2000-01-16', 'Guajira', 'Riohacha');
insert into cliente values (null, 13423584, 'Mauricio', 'Macri', '1990-03-23', 'Buenos Aires', 'Ciudad Capital');

INSERT INTO plan VALUES(null, 10, 800, 0, 21);
INSERT INTO plan VALUES(null, 100, 200, 0, 22);
INSERT INTO plan VALUES(null, 500, 4000, 5, 23);
INSERT INTO plan VALUES(null, 1000, 8000, 5, 24);
INSERT INTO plan VALUES(null, 10000, 32000, 10, 25);
​
select id, nombre, dni
from cliente
where id >= 10;
​
select precio,velocidad_mbps
from plan 
where velocidad_mbps
between 100 and 1000;
​
SELECT AVG(precio) as 'Precio promedio' 
FROM plan;
​
select min(velocidad_mbps) 
from plan;
​
select * 
from cliente
where provincia like "bu%" 
Order By dni desc;
​
select year(fecha_nacimiento), nombre, apellido 
from cliente;
​
SELECT COUNT(id) as 'Total clientes' 
FROM cliente;
​
UPDATE plan SET precio = 1500 
WHERE velocidad_mbps = 500;
​
DELETE FROM plan 
WHERE id_cliente = (select id from cliente where nombre = 'Maury');
​
Delete From cliente
where nombre = 'Maury';
​
select * from plan;
​
select * from cliente;

