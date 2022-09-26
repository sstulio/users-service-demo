FROM golang:1.16-alpine as build

WORKDIR /server

COPY go.* ./
COPY cmd ./cmd
COPY internal ./internal
COPY vendor ./vendor

RUN mkdir -p build && cd ./build && \
    go build -o application -mod vendor ../cmd

FROM alpine:3

COPY --from=build /server/build/* ./

EXPOSE 8080

ENTRYPOINT ["./application"]