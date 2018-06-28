FROM ataraev/golang-alpine-git:latest
MAINTAINER Jack Frost <j4qfrost@gmail.com>

COPY src /go/src/app
WORKDIR /go/src/app

VOLUME /go/src/app

RUN apk add --update --no-cache gcc && \
	apk add --update --no-cache g++

RUN go get -d -v ./... && go install -v ./... && go build

CMD ["app"]