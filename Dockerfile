FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /lwwset

WORKDIR /lwwset

COPY . .

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/lwwset


FROM scratch

COPY --from=builder /go/bin/lwwset /go/bin/lwwset

ENTRYPOINT ["/go/bin/lwwset"]

EXPOSE 8080