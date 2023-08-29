FROM golang:1.21.0-alpine

RUN apk update && apk upgrade && apk add make 

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN make install

COPY . ./

RUN make build -B

EXPOSE 8080

CMD ["make", "run"]