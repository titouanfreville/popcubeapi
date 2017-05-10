FROM registry.popcube.xyz:5000/go:base
MAINTAINER Clement LE CORRE <clement@le-corre.eu>
ENV TERM xterm-256color
ENV GOCOPYPATH /go/src/github.com/titouanfreville/popcubeapi
ENV MYSQL_DATABASE popcube_test
EXPOSE 8080
WORKDIR $GOCOPYPATH

ENTRYPOINT ["goconvey","-host","0.0.0.0","-cover"]
CMD ["-excludedDirs","github.com,golang.org,gopkg.in,vendor"]
