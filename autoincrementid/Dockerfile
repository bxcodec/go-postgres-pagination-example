# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.14.0-stretch as builder

RUN apt --yes --force-yes update && apt --yes --force-yes upgrade && \
    apt --yes --force-yes install git \
    make openssh-client

WORKDIR /go_modules/auto-increment-id

COPY . .
RUN make auto-increment

## Distribution
FROM debian:stretch-slim

RUN apt update && apt upgrade -y && \
    apt install -y ca-certificates tzdata && \
    mkdir /app && mkdir auto-increment-id

WORKDIR /auto-increment-id

EXPOSE 9090

COPY --from=builder /go_modules/auto-increment-id/aiengine /app

CMD /app/aiengine
