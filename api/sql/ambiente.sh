#!/bin/bash
docker volume create mysql-volume
docker run --name container-mysql -p3306:3306 -v mysql-volume:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -d mysql/mysql-server:latest
docker exec -it container-mysql bash

mysql -u root -p
create user 'golang'@'%' identified by 'golang';
update mysql.user set host='%' where user='root';
update mysql.user set host='%' where user='golang';
grant all privileges on *.* to 'golang'@'%';
flush privileges;