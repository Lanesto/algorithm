# Image name to be used
$image = "algorithm_go:local"
 
# Embedded dockerfile
$dockerfile = `
'FROM golang:1.15

USER root
VOLUME /workspace

RUN apt update && \
    apt install -y \
        vim \
        strace \
        dos2unix

RUN go get -v golang.org/x/lint/golint

WORKDIR /workspace

CMD ["/bin/bash"]'

# Create temporary dockerfile
$temp = New-TemporaryFile
Set-Content -Path $temp -Value $dockerfile

# Build image with temporary dockerfile
docker image build -f $temp -t $image .
echo ""

# Run
docker container run -it --rm `
    -v ${PWD}:/workspace `
    -w /workspace `
    $image /bin/bash

# Exited container; give user some time to see outputs
pause