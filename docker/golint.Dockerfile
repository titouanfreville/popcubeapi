FROM cl3m3nt/golinter
LABEL MAINTAINER "clement@le-corre.eu"
ENV TERM xterm-256color
COPY api /go/src/api
COPY models /go/src/models
COPY utils /go/src/utils
COPY datastores /go/src/datastores
COPY main.go /go/src/main.go
