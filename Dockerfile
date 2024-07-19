###
# Build Stage
###
FROM golang:1.22.2-alpine3.19 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/apollo


###
# Release Stage
###
FROM golang:1.22.2-alpine3.19 AS release-stage

WORKDIR /app

RUN apk update && apk add --no-cache curl

COPY --from=build-stage /usr/local/bin/apollo /usr/local/bin/apollo

EXPOSE 3400/tcp
# EXPOSE 3400/udp
USER root

HEALTHCHECK --interval=15s --timeout=10s --start-period=30s --retries=3 \
  CMD curl -f http://localhost:3300/api/health/v1 || exit 1

CMD [ "apollo" ]
