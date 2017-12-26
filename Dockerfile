FROM ataraev/golang-alpine-git:latest
MAINTAINER Jack Frost <j4qfrost@gmail.com>

COPY src /go/src/app
WORKDIR /go/src/app

VOLUME /go/src/app

RUN apk add --update --no-cache gcc && \
	apk add --update --no-cache g++ && \
	go-wrapper download && go-wrapper install
CMD ["go-wrapper", "run"]