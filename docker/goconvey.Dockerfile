FROM registry.le-corre.eu:5000/go:base
MAINTAINER Clement LE CORRE <clement@le-corre.eu>
ENV TERM xterm-256color
ENV GOCOPYPATH /go/src/github.com/titouanfreville/popcubeapi

EXPOSE 8080
WORKDIR $GOCOPYPATH

ENTRYPOINT ["goconvey","-host","0.0.0.0","-cover"]
CMD ["-excludedDirs","github.com,golang.org,gopkg.in,vendor"]
