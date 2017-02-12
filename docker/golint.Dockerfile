FROM cl3m3nt/golinter
ENV TERM xterm-256color
MAINTAINER Clement LE CORRE <clement@le-corre.eu>
COPY api /go/src/api
COPY models /go/src/models
COPY utils /go/src/utils
COPY datastores /go/src/datastores
COPY main.go /go/src/main.go
