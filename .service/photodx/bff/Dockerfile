# Build bff
FROM golang:1.22-alpine AS build2

RUN apk add --no-cache make

WORKDIR /app

COPY . ./
RUN make photodx/bff-build

# Image
FROM alpine:latest

WORKDIR /app
COPY --from=build2 /app/photodx-bff.out .

ENTRYPOINT ["/app/photodx-bff.out"]