FROM golang:1.11-alpine as builder
RUN apk add bash ca-certificates git gcc g++ libc-dev
RUN mkdir /build
ADD . /build/
WORKDIR /build
ENV GO111MODULE=on
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
ENTRYPOINT ["./main"]