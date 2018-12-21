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

# Golang TensorFlow Protocol Buffers

## Compilation 

GOTFPB adds Golang-specific optimizations. These optimizations depend on a 3rd party fork of Google's protoc-gen code which improves the quality of code generated by the protoc tool. A useful generator script has been made to simplify compilation even further. Install third-party dependencies and then run generate in the repo root. 

```bash
# generate necessary code files
go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
go generate ./...
```
