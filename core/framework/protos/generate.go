//go:generate go run generate.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func main() {
	if err := generateProtos(); err != nil {
		log.Error().Err(err).Msg("compilation error")
	}
	log.Info().Msg("successfully compiled all protocol buffers")
}

func generateProtos() error {
	files, err := filepath.Glob("*.proto")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	args := []string{
		"-I=protos",
		"--go_out=.",
	}
	args = append(args, files...)
	cmd := exec.Command("protoc", args...)
	fmt.Println(cmd)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
