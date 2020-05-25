# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.14.0-stretch as builder

RUN apt --yes --force-yes update && apt --yes --force-yes upgrade && \
    apt --yes --force-yes install git \
    make openssh-client

WORKDIR /go_modules/page-number

COPY . .
RUN make page-number

## Distribution
FROM debian:stretch-slim

RUN apt update && apt upgrade -y && \
    apt install -y ca-certificates tzdata && \
    mkdir /app && mkdir page-number

WORKDIR /page-number

EXPOSE 9090

COPY --from=builder /go_modules/page-number/pnengine /app

CMD /app/pnengine
