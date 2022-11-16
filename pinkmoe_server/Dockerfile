FROM golang:latest

WORKDIR /pinkmoe_server
COPY . .

ENV TIME_ZONE Asia/Shanghai

RUN apt-get install -y tzdata \
    && ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata \
    && go generate && go env && go build -o server .

EXPOSE 9527

ENTRYPOINT ./server config.yaml