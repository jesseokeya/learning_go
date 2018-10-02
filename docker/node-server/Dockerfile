FROM node:10-alpine

WORKDIR /usr/app

COPY ./package*.json ./

ENV PORT=3002

RUN npm install

COPY ./ ./

CMD ["npm", "start"]