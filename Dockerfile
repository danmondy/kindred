FROM golang:1.22.0-alpine3.19

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /kindred

EXPOSE 8080

CMD [ "/kindred" ]