<p align="center">
<img width="200" alt="TensorTask Logo" src="https://storage.googleapis.com/tensortask-static/tensortask_transparent.png">
</p>

[![GoDoc][1]][2] [![Go Report Card][3]][4] [![Keybase Chat][5]][6] [![Cloud Build][7]][8]

[1]: https://godoc.org/github.com/tensortask/gotfpb?status.svg
[2]: https://godoc.org/github.com/tensortask/gotfpb
[3]: https://goreportcard.com/badge/github.com/tensortask/gotfpb
[4]: https://goreportcard.com/report/github.com/tensortask/gotfpb
[5]: https://img.shields.io/badge/keybase%20chat-tensortask.public-blue.svg
[6]: https://keybase.io/team/tensortask.public
[7]: https://storage.googleapis.com/tensortask-static/build/gotfpb.svg
[8]: https://github.com/sbsends/cloud-build-badge

# 📜 GoTFPB: Curated Golang TensorFlow Protocol Buffers
GoTFPB provides compiled golang code for tensorflow protocol buffers. Instead of compiling the protocol buffers locally, you can simply import the golang code like any other golang package.

Note: The build is failing because no unit tests have been written yet.

## Package Layout

Each package has two sub-directories. The gen directory contains the generated Golang source code. The protos directory contains the protocol buffer definitions (proto3).

* **core**: fundamental protocol buffers used at the core of TensorFlow.
  * **framework**: basic absolutely vital tensors used in TensorFlow (graph, tensor, etcetera). 

## Imports

* **core**:
  * **framework**: `github.com/tensortask/gotfpb/core/framework`

## Compilation

Local compilation is not required, but if you wish to recompile the protos we reccommend using the [supplied docker image](https://github.com/tensortask/go-proto-compiler).

Build the docker image by running the following commands:

```bash
git clone https://github.com/tensortask/go-proto-compiler
cd go-proto-compiler
docker build . -t go_proto_compiler
```

Compile the protos:
```bash
# first move to the gotfpb dir
cd $GOPATH/src/github.com/tensortask/gotfpb

# unix-based platforms
bash compile.sh

# or run the docker command (non-unix)
docker run -v $GOPATH/src/github.com/tensortask/gotfpb:/go/src/github.com/tensortask/gotfpb go_proto_compiler /bin/bash -c "export GO111MODULE=on; cd src/github.com/tensortask/gotfpb; go generate ./..."
```

