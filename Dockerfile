FROM ubuntu:20.04

LABEL version="1.0"
LABEL maintainer="sirily11"

ENV DEBIAN_FRONTEND=noninteractive

# Install Go
RUN apt update
RUN apt-get -y install software-properties-common
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt -y install golang-go

# Install bash
RUN apt-get install bash

# Install make
RUN apt-get install -y make

# Install geth
COPY . /build/etd
WORKDIR /build/etd
RUN make geth

# Export geth command
ENV PATH="/build/go-ethereum/build/bin:${PATH}"
RUN geth --help