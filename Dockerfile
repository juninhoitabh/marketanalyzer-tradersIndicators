FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install build-essential sudo protobuf-compiler -y && \
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get github.com/spf13/cobra/cobra && \
    wget https://github.com/ktr0731/evans/releases/download/0.10.0/evans_linux_amd64.tar.gz && \
    tar -xzvf evans_linux_amd64.tar.gz && \
    mv evans ../bin && rm -f evans_linux_amd64.tar.gz && \
    wget http://prdownloads.sourceforge.net/ta-lib/ta-lib-0.4.0-src.tar.gz && \
    tar -xzf ta-lib-0.4.0-src.tar.gz && \
    cd ta-lib -0.4.0-src.tar.gz && \
    ./configure --prefix=/usr LDFLAGS="-lm" && \
    make && make install

CMD ["tail", "-f", "/dev/null"]