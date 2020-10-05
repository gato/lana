FROM ubuntu:latest as build
RUN apt-get update
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get -y upgrade
RUN apt-get -y install golang ca-certificates
WORKDIR /api/
COPY go.mod go.sum main.go /api/
COPY checkout /api/checkout
COPY merchandise /api/merchandise
RUN go build

FROM ubuntu:latest
RUN apt-get update && apt-get -y upgrade && apt-get install -y libc6 ca-certificates tzdata && rm -rf /var/lib/apt/lists/*
WORKDIR /api/
COPY --from=build /api/lana /api/lana
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["/api/lana"]
