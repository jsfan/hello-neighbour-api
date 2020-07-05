#!/bin/sh

# if there is no Postgres certificate, we create a self-signed one
if [[ ! -e /etc/postgres/ssl/postgres.key || ! -e /etc/postgres/ssl/postgres.crt ]]
then
    openssl req -new -newkey rsa:4096 -days 3650 -nodes -x509 -keyout /etc/postgres/ssl/postgres.key -out etc/postgres/ssl/postgres.crt -subj '/CN=postgres/'
    chown -R postgres:postgres /etc/postgres
fi

/usr/local/bin/docker-entrypoint.sh -c 'config_file=/etc/postgres/postgresql.conf'
