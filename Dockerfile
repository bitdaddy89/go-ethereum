FROM ubuntu:20.04

LABEL version="1.0"
LABEL maintainer="sirily11"

ENV DEBIAN_FRONTEND=noninteractive

WORKDIR /home/build

# Install Go
RUN apt update
RUN apt-get -y install software-properties-common
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt -y install golang-go

# Install bash
RUN apt-get install bash

# Install Python3
RUN apt-get install -y python3
RUN apt-get install -y python3-pip

# Install geth
COPY . /build/go-ethereum
WORKDIR /build/go-ethereum
RUN make geth

# Export geth command
ENV PATH="/build/go-ethereum/build/bin:${PATH}"
RUN geth --help