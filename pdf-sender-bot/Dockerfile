FROM golang:1.20-buster as builder

RUN apt update -y && apt install make -y
WORKDIR /opt
COPY . .
RUN make build

FROM ubuntu:jammy
RUN apt update -y && apt install wkhtmltopdf -y && rm -rf /var/lib/{apt,dpkg}/*
RUN adduser --disabled-password --gecos '' bot
COPY --from=builder --chown=bot /opt/bin/app /opt/app
COPY --from=builder --chown=bot /opt/config/template.html /opt/config/template.html
WORKDIR /opt
USER bot
CMD ["./app"]