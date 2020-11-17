FROM golang:1.15

USER root
VOLUME /workspace

RUN apt update && \
    apt install -y \
        vim \
        strace \
        dos2unix

RUN go get -v golang.org/x/lint/golint

WORKDIR /workspace

CMD ["/bin/bash"]
