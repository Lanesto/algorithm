FROM golang:1.15

# Install python3
RUN apt update && \ 
    apt install -y --no-install-recommends \
        python3 \
        python3-pip \
        python3-dev \
        python3-setuptools \
    && rm -rf /var/lib/apt/lists/*

# Install python requirements
ADD .devcontainer/requirements.txt /
RUN pip3 install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir -r /requirements.txt
