FROM node:19 as ts

COPY src/ts /src
WORKDIR /src/vue-project
RUN npm install && npm run build

FROM golang:1.22 AS go

COPY src/go /src
COPY --from=ts /src/vue-project/dist /src/go-project/static
WORKDIR /src/go-project
RUN go build

FROM alpine:latest

RUN apk add libc6-compat
COPY --from=go /src/go-project/go-project /app/go-project

EXPOSE 3000
ENTRYPOINT ["/app/go-project"]