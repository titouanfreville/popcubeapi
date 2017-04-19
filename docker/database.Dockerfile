FROM mariadb:10.1
LABEL MAINTAINER "clement@le-corre.eu"
COPY scripts/init_values.sql /docker-entrypoint-initdb.d/init_values.sql
COPY scripts/init.sql /docker-entrypoint-initdb.d/init.sql
