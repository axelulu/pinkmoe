FROM node:16.16.0

WORKDIR /pinkmoe_index
COPY . .

RUN npm install npm@latest -g && npm install pnpm -g && pnpm i && pnpm build

EXPOSE 3000

ENTRYPOINT pnpm start