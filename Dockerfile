# cf. https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076
# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.21.0

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD air
