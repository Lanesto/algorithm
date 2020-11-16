FROM golang:1.15

VOLUME /workspace

RUN apt update && \
    apt install -y \
        vim \
        rsyslog \
        strace \
        sqlite3 \
        graphviz

RUN go get -v golang.org/x/tools/cmd/godoc

WORKDIR /workspace
    
COPY ./docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["/bin/bash"]
