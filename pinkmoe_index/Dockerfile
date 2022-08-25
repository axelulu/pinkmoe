FROM node:16.16.0

WORKDIR /pinkmoe_index
COPY . .

RUN npm install npm@latest -g && npm install pm2 -g && npm install pnpm -g && pnpm i && pnpm build && pm2 start ecosystem.config.js
