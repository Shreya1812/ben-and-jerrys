FROM golang:1.13

RUN apt-get update || exit 0
RUN apt-get upgrade -y

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o  main ./cmd/server

EXPOSE 9000

RUN bash ./deployments/pre.sh

CMD ./main > out.log

