FROM surnet/alpine-wkhtmltopdf:3.17.0-0.12.6-full as wkhtmltopdf
FROM golang:1.20-alpine3.17 as builder

RUN apk add --no-cache make
WORKDIR /opt
COPY . .
RUN make build

FROM alpine:3.17
RUN addgroup -g 65535 bot && \
    adduser --shell /sbin/nologin --disabled-password \
    --no-create-home --uid 65535 --ingroup bot bot
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/libwkhtmltox.so /bin/
COPY --from=builder --chown=bot /opt/bin/app /opt/app
WORKDIR /opt
USER bot
CMD ["./app"]