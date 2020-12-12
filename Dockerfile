FROM node:alpine
RUN apk add python3
WORKDIR /app
COPY nuxt.config.js yarn.lock package.json LICENSE ./
RUN yarn install
COPY pages ./pages
COPY layouts ./layouts
COPY assets ./assets
COPY components ./components
COPY plugins ./plugins
CMD yarn dev
