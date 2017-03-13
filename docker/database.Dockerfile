FROM mariadb:10.1
MAINTAINER Clement LE CORRE <clement@le-corre.eu>
COPY scripts/*.sql /docker-entrypoint-initdb.d/
