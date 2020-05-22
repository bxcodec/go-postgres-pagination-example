# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.14.0-stretch as builder

RUN apt --yes --force-yes update && apt --yes --force-yes upgrade && \
    apt --yes --force-yes install git \
    make openssh-client

WORKDIR /go_modules/offset-limit

COPY . .
RUN make offset-limit

## Distribution
FROM debian:stretch-slim

RUN apt update && apt upgrade -y && \
    apt install -y ca-certificates tzdata && \
    mkdir /app && mkdir offset-limit

WORKDIR /offset-limit

EXPOSE 9090

COPY --from=builder /go_modules/offset-limit/olengine /app

CMD /app/olengine
