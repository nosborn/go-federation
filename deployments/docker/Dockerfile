FROM alpine:3.22.0

RUN apk add --no-cache supervisor=4.2.5-r5 \
    && adduser -h /app -s /sbin/nologin -D -H -u 1000 federation

COPY --chmod=0600 supervisord.conf /etc/supervisord.conf

WORKDIR /app
ENTRYPOINT ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf", "-n"]
