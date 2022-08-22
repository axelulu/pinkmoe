FROM golang:latest

WORKDIR /pinkmoe_server
COPY . .

RUN apt-get update -y && apt-get upgrade -y && apt-get install libvips-dev -y && go generate && go env && go build -o server .

EXPOSE 9527

ENTRYPOINT ./server config.yaml