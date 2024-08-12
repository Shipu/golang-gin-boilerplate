ARG GO_VERSION=1.22

FROM golang:${GO_VERSION} as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

COPY . .

CMD ["air"]

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

FROM golang:${GO_VERSION}-alpine as prod

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY --from=builder /api/docs ./docs
COPY --from=builder /api/app .

EXPOSE 8081

ENTRYPOINT ["./app"]