FROM postgres:11.4-alpine

ADD ./payment_incremental_id.sql /docker-entrypoint-initdb.d