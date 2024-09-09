FROM nvidia/cuda:11.1.1-devel-ubuntu20.04

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -y \
      build-essential git wget pkg-config && \
    rm -rf /var/lib/apt/lists/*

ENV GO_VERSION 1.21.3
RUN mkdir /build && cd /build && \
    wget --progress=bar:force https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /build
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin