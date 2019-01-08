# To gurantee reproducible builds we use a docker image: https://github.com/tensortask/go-proto-compiler.
# This compile scripts requires that the image is built prior to running the script.
# git clone https://github.com/tensortask/go-proto-compiler
# cd go-proto-compiler
# docker build . -t go_proto_compiler
docker run -v $GOPATH/src/github.com/tensortask/gotfpb:/go/src/github.com/tensortask/gotfpb go_proto_compiler /bin/bash -c "export GO111MODULE=on; cd src/github.com/tensortask/gotfpb; go generate ./..."