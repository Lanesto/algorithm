$image = "golang:1.15-local"

docker image build -t $image .
echo ""

docker container run -it --rm `
    -v ${PWD}:/workspace `
    -w /workspace `
    $image /bin/bash

pause