# ===== Build stage of frontend =====
#FROM node:20-alpine3.18 as build-frontend
#
#WORKDIR /app
#
#COPY . .
#
#WORKDIR /app/slimdough
#
#RUN npm ci
#RUN npm run build

# ===== Build stage of application =====
FROM golang:1.21-alpine3.19 AS build-server

WORKDIR /app

COPY go.* ./
RUN go mod download # && go mod verify

COPY . .

RUN go build -v -o /pork-go

# ===== Deploy app bin into a lean image =====
FROM alpine:3.19.0 AS build-release-stage

WORKDIR /

#RUN apt-get -y update
#RUN apt-get -y install ca-certificates

#COPY --from=build-frontend /app/slimdough/dist /dist
COPY --from=build-server /pork-go /pork-go

EXPOSE 8000

CMD ["/pork-go"]
