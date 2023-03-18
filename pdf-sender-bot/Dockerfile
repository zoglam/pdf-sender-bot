FROM golang:1.20-alpine3.17 as builder

RUN apk add --no-cache make
WORKDIR /opt
COPY . .
RUN make build

FROM surnet/alpine-wkhtmltopdf:3.17.0-0.12.6-full
RUN addgroup -g 65535 bot && \
    adduser --shell /sbin/nologin --disabled-password \
    --no-create-home --uid 65535 --ingroup bot bot
COPY --from=builder --chown=bot /opt/bin/app /opt/app
WORKDIR /opt
USER bot
ENTRYPOINT [""]
CMD ["./app"]