FROM mariadb:10.1
MAINTAINER Clement LE CORRE <clement@le-corre.eu>
COPY docker-entrypoint-initdb.d /docker-entrypoint-initdb.d/
RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/

ENV ORG_ORGANISATIONNAME=None \
    ORG_DESCRIPTION=None \
    ORG_AVATAR=None \
    ORG_DOMAIN=None
