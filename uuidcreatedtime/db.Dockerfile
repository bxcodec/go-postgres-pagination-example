FROM postgres:11.4-alpine

ADD ./payment_with_uuid.sql /docker-entrypoint-initdb.d