FROM golang:1.15

# Install python
RUN apt update && \ 
    apt install -y --no-install-recommends \
        python3 \
        python3-pip \
        python3-dev \
        python3-setuptools \
    && rm -rf /var/lib/apt/lists/*

# Install python requirements
RUN pip3 install --no-cache-dir \
        flake8 \
        pycodestyle_magic \
        jupyter

# Install go tools
RUN go get -v \
        golang.org/x/tools/cmd/godoc \
        golang.org/x/tools/cmd/goimports \
        golang.org/x/lint/golint \
        golang.org/x/tools/gopls \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/stamblerre/gocode \
        github.com/rogpeppe/godef \
        github.com/sqs/goreturns

RUN echo 'alias pip=pip3' >> ~/.bashrc
