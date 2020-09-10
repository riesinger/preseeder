FROM golang:alpine as builder

RUN mkdir /app && apk update && apk add git
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o preseeder ./cmd/preseeder

FROM scratch as runner

WORKDIR /app
COPY --from=builder /app/preseeder /app/preseeder
ENV BASE_DIR /data
EXPOSE 3000

CMD [ "/app/preseeder" ]
