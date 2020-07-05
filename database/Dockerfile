FROM postgres:12-alpine

RUN apk add openssl

RUN mkdir -p /etc/postgres
COPY postgresql.conf /etc/postgres

COPY entrypoint.sh /
RUN chmod 755 /entrypoint.sh

VOLUME /etc/postgres/ssl

RUN chown postgres:postgres /etc/postgres/ssl

ENTRYPOINT [ "/entrypoint.sh" ]