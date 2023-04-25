FROM golang:1.20-alpine as builder
WORKDIR /

COPY . .
RUN go mod download

RUN go build -o serfer .

FROM scratch
WORKDIR /bin
COPY --from=builder /serfer /bin

CMD [ "/bin/serfer" ]