FROM golang:1.20-alpine as builder
WORKDIR /

COPY . .
RUN go mod download

RUN go build -o distcache .

FROM scratch
WORKDIR /bin
COPY --from=builder /distcache /bin

CMD [ "/bin/distcache" ]