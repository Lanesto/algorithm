FROM golang:1.15

USER root
VOLUME /workspace

RUN apt update && \
    apt install -y vim strace

WORKDIR /workspace

CMD ["/bin/bash"]
