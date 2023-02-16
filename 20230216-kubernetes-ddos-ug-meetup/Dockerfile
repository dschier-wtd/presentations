FROM docker.io/library/node:16 as build

WORKDIR /app

COPY ember-sample/. /app/

RUN npm install
RUN ./node_modules/ember-cli/bin/ember build

COPY --from=build --chown=33:33 /app/dist/ /usr/local/apache2/htdocs/
