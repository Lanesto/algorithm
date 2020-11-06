$image = "algorithm_go:local"

docker image build -f "go.Dockerfile" -t $image .
echo ""

docker container run -it --rm `
    -v ${PWD}:/workspace `
    -w /workspace `
    $image /bin/bash

pause