FROM golang:1.14.2-alpine as base

RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -installsuffix cgo -o app cmd/main.go

FROM scratch

COPY --from=base /etc/ssl/certs /etc/ssl/certs
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /build/app /bin/

ENTRYPOINT [ "/bin/app" ]