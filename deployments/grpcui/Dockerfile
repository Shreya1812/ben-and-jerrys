FROM golang:1.13

RUN apt-get update || exit 0
RUN apt-get upgrade -y

ENV GO111MODULE=on

WORKDIR /

COPY . .

EXPOSE 9999
EXPOSE 9991


RUN go get github.com/fullstorydev/grpcui
RUN go install github.com/fullstorydev/grpcui/cmd/grpcui

CMD  bash ./start.sh
