FROM golang:1.15

WORKDIR /work

COPY * ./

RUN go build

ENTRYPOINT [ "/work/run.sh" ]