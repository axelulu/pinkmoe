FROM golang:alpine

WORKDIR /pinkmoe_server
COPY . .

RUN go generate && go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /pinkmoe_server

COPY --from=0 /pinkmoe_server ./

EXPOSE 9527

ENTRYPOINT ./server config.yaml