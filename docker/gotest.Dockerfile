FROM registry.le-corre.eu:5000/go:base
MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

ENV TERM xterm-256color
ENV WATCHING 0
ENV TERM xterm-256color
ENV GOCOPYPATH go/src/github.com/titouanfreville/popcubeapi
ENV GOSU_VERSION 1.9

COPY scripts/wait-for-it.sh /bin/waitforit
COPY scripts/go_test_entrypoint.sh /bin/entrypoint

WORKDIR /go/src

RUN set -x \
    && apk add --no-cache --virtual .gosu-deps \
        dpkg \
        gnupg \
        openssl \
    && dpkgArch="$(dpkg --print-architecture | awk -F- '{ print $NF }')" \
    && wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch" \
    && wget -O /usr/local/bin/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch.asc" \
    && export GNUPGHOME="$(mktemp -d)" \
    && gpg --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 \
    && gpg --batch --verify /usr/local/bin/gosu.asc /usr/local/bin/gosu \
    && rm -r "$GNUPGHOME" /usr/local/bin/gosu.asc \
    && chmod +x /usr/local/bin/gosu \
    && gosu nobody true \
    && apk del .gosu-deps

ENTRYPOINT entrypoint /$GOCOPYPATH $WATCHING
