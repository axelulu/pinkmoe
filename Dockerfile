FROM node:16.16.0

WORKDIR /
COPY . .

RUN cd ./pinkmoe_index && npm install npm@latest -g && yarn && yarn build

RUN cd ../pinkmoe_admin && npm install npm@latest -g && yarn && yarn build

FROM nginx:alpine
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

# index
COPY ./pinkmoe_index/nginx/pinkmoe.conf /etc/nginx/conf.d/pinkmoe.conf
COPY ./pinkmoe_index/nginx/pinkmoe_ssl.conf /etc/nginx/conf.d/pinkmoe_ssl.conf
COPY ./pinkmoe_index/nginx/ssl/pinkmoe.key /etc/nginx/pinkmoe.key
COPY ./pinkmoe_index/nginx/ssl/pinkmoe.crt /etc/nginx/pinkmoe.crt
# admin
COPY ./pinkmoe_admin/nginx/admin_pinkmoe.conf /etc/nginx/conf.d/admin_pinkmoe.conf
COPY ./pinkmoe_admin/nginx/admin_pinkmoe_ssl.conf /etc/nginx/conf.d/admin_pinkmoe_ssl.conf
COPY ./pinkmoe_admin/nginx/ssl/admin.pinkmoe.key /etc/nginx/admin.pinkmoe.key
COPY ./pinkmoe_admin/nginx/ssl/admin.pinkmoe.crt /etc/nginx/admin.pinkmoe.crt
RUN rm -rf /etc/nginx/conf.d/default.conf
RUN mkdir /usr/share/nginx/html/index
RUN mkdir /usr/share/nginx/html/admin
COPY --from=0 /pinkmoe_index/dist /usr/share/nginx/html/index
COPY --from=0 /pinkmoe_admin/dist /usr/share/nginx/html/admin
RUN cat /etc/nginx/nginx.conf
# index
RUN cat /etc/nginx/conf.d/pinkmoe.conf
RUN cat /etc/nginx/conf.d/pinkmoe_ssl.conf
# admin
RUN cat /etc/nginx/conf.d/admin_pinkmoe.conf
RUN cat /etc/nginx/conf.d/admin_pinkmoe_ssl.conf
RUN ls -al /usr/share/nginx/html