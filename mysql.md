### 1. Pull mysql images to localhost

    docker pull mysql:tag

### 2. Run mysql

    docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:tag
    docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:tag
    
### 3. Enter mysql container

    docker exec -it container_name|container_id bash
    
### 4. See mysql logs

    docker logs container_name|container_id
    
### 5. Use custom configuration parameters

    docker run --name mysql -v /root/mysql:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=root -d mysql:tag
    
### 6. Mount the container data location and host location to ensure data security

    docker run --name mysql -v /root/mysql/data:/var/lib/mysql -v /root/mysql/config.d:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:tag
    
### 7. Access via other clients

### 8. Back up the mysql database as a sql file

    docker exec mysql sh -c 'exec mysqldump --all-databases -uroot -p"$MYSQL_ROOT_PASSWORD"' > /root/all-databases.sql    ==>export all data
    docker exec mysql sh -c 'exec mysqldump --databases -uroot -p"$MYSQL_ROOT_PASSWORD"' > /root/all-databases.sql    ==>export specific data
    docker exec mysql sh -c 'exec mysqldump --no-data --databases -uroot -p"$MYSQL_ROOT_PASSWORD"' > /root/all-databases.sql    
    
### 9. Execute sql file into mysql

    docker exec -i mysql sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD"' < /root/xxx.sql
