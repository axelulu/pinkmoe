FROM node:16.16.0

WORKDIR /
COPY . .

RUN export NODE_OPTIONS=--max-old-space-size=8192 && cd ./pinkmoe_admin && npm install npm@latest -g && npm i && npm run build

FROM centos:centos7

WORKDIR /
COPY . .

RUN yum update -y && curl -fsSL https://rpm.nodesource.com/setup_16.x | bash - && yum install -y nodejs && node --version && npm --version
RUN rpm --rebuilddb;yum install make wget tar gzip passwd openssh-server gcc pcre-devel openssl-devel net-tools vim -y
RUN wget -P /tmp/ http://nginx.org/download/nginx-1.23.1.tar.gz
RUN cd /tmp/;tar xzf nginx-1.23.1.tar.gz;cd nginx-1.23.1;sed -i -e 's/1.23.1//g' -e 's/nginx\//WS/g' -e 's/"NGINX"/"WS"/g' src/core/nginx.h
RUN cd /tmp/nginx-1.23.1;./configure --prefix=/usr/local/nginx --with-http_v2_module --with-http_stub_status_module --with-http_ssl_module;make;make install

RUN export NODE_OPTIONS=--max-old-space-size=8192 && cd /pinkmoe_index && npm install npm@latest -g && npm install pnpm -g && pnpm i && pnpm build

RUN mkdir /usr/local/nginx/conf.d
RUN rm -rf /usr/local/nginx/conf/nginx.conf
COPY ./default.conf /usr/local/nginx/conf/nginx.conf
# index
COPY ./pinkmoe_index/nginx/pinkmoe.conf /usr/local/nginx/conf/conf.d/pinkmoe.conf
COPY ./pinkmoe_index/nginx/pinkmoe_ssl.conf /usr/local/nginx/conf/conf.d/pinkmoe_ssl.conf
COPY ./pinkmoe_index/nginx/ssl/pinkmoe.key /usr/local/nginx/pinkmoe.key
COPY ./pinkmoe_index/nginx/ssl/pinkmoe.crt /usr/local/nginx/pinkmoe.crt
# admin
COPY ./pinkmoe_admin/nginx/admin_pinkmoe.conf /usr/local/nginx/conf/conf.d/admin_pinkmoe.conf
COPY ./pinkmoe_admin/nginx/admin_pinkmoe_ssl.conf /usr/local/nginx/conf/conf.d/admin_pinkmoe_ssl.conf
COPY ./pinkmoe_admin/nginx/ssl/admin.pinkmoe.key /usr/local/nginx/admin.pinkmoe.key
COPY ./pinkmoe_admin/nginx/ssl/admin.pinkmoe.crt /usr/local/nginx/admin.pinkmoe.crt
RUN rm -rf /usr/local/nginx/conf/conf.d/default.conf
# RUN mkdir /usr/share/nginx/html/index
RUN mkdir /usr/share/nginx/html/admin -p
COPY --from=0 ./pinkmoe_admin/dist /usr/share/nginx/html/admin
RUN cat /usr/local/nginx/conf/nginx.conf
# index
RUN cat /usr/local/nginx/conf/conf.d/pinkmoe.conf
RUN cat /usr/local/nginx/conf/conf.d/pinkmoe_ssl.conf
# admin
RUN cat /usr/local/nginx/conf/conf.d/admin_pinkmoe.conf
RUN cat /usr/local/nginx/conf/conf.d/admin_pinkmoe_ssl.conf
RUN ls -al /usr/share/nginx/html

EXPOSE 443
EXPOSE 80
CMD /usr/local/nginx/sbin/nginx;node /pinkmoe_index/.output/server/index.mjs;/usr/sbin/sshd -D