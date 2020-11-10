FROM golang:1.15

USER root
VOLUME /workspace

RUN apt update && \
    apt install -y vim rsyslog strace sqlite3

WORKDIR /workspace
    
COPY ./docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["/bin/bash"]
