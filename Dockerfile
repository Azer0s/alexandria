FROM golang:latest as build

WORKDIR /go/src/github.com/Azer0s/alexandria
COPY . .

RUN go version
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/Azer0s/alexandria/app .
CMD ["./app"]
