//go:generate go run generate.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

var (
	goPath = os.Getenv("GOPATH")
)

func main() {
	if err := generateProtos(); err != nil {
		log.Error().Err(err).Msg("Compilation error.")
	}
}

func generateProtos() error {
	files, err := filepath.Glob("protos/*")
	if err != nil {
		log.Error().Err(err).Msg("could not read directory")
	}
	protofiles, err := splitPath(files)
	if err != nil {
		log.Error().Err(err).Msg("paths could not be split")
	}

	for _, path := range protofiles {
		args := []string{
			fmt.Sprintf("-I=%s", filepath.Join(goPath, "src")),
			fmt.Sprintf("-I=%s", filepath.Join(goPath, "src", "github.com", "gogo", "protobuf", "protobuf")),
			fmt.Sprintf("-I=%s", filepath.Join(goPath, "src", "github.com", "tensortask", "gotfpb", "core", "framework", "protos")),
			fmt.Sprintf("--proto_path=%s", filepath.Join(goPath, "src", "github.com")),
			fmt.Sprintf("--proto_path=%s", filepath.Join(goPath, "src", "github.com", "tensortask/gotfpb/core/framework/protos")),
			"--gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:gen",
			path,
		}

		cmd := exec.Command("protoc", args...)
		fmt.Println("\n")
		fmt.Println("\n")
		fmt.Println(cmd)
		fmt.Println("\n")
		fmt.Println("\n")
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	log.Info().Msg("Successfully compiled TTP protocol buffers.")
	return nil
}

func splitPath(paths []string) ([]string, error) {
	var files []string
	for _, p := range paths {
		_, file := filepath.Split(p)
		files = append(files, file)
	}
	return files, nil
}
