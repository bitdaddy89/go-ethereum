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

# Install getd
COPY . /build/etd
WORKDIR /build/etd
RUN make getd

# Export getd command
ENV PATH="/build/etd/build/bin:${PATH}"
RUN getd --help