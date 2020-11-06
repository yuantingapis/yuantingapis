// Command gen-yuantingapis auto generates api clients for all supported languages.

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	wd = flag.String("path", "", "Root directory for yuantingapis project")
)

func genProtoGo(wd, protoPath string) {
	p := filepath.Dir(wd)

	files, err := filepath.Glob(fmt.Sprintf("%v/yuanting/yt/v1/*.proto", wd))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
		return
	}
	args := []string{
		"-I", wd,
		"-I", protoPath,
		fmt.Sprintf("--go_out=%v/go-genproto", p),
		"--go_opt=paths=source_relative",
		fmt.Sprintf("--go-grpc_out=%v/go-genproto", p),
		"--go-grpc_opt=paths=source_relative",
	}
	args = append(args, files...)
	_, err = exec.Command("protoc", args...).CombinedOutput()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func genClientGo(wd, protoPath string) {
	p := filepath.Dir(wd)

	files, err := filepath.Glob(fmt.Sprintf("%v/yuanting/yt/v1/*.proto", wd))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
		return
	}
	args := []string{
		"-I", wd,
		"-I", protoPath,
		fmt.Sprintf("--go_gapic_out=%v/go-client", p),
		"--go_gapic_opt", "go-gapic-package=github.com/yuantingapis/go-client/yuanting/yt/v1;yt",
		"--go_gapic_opt", "module=github.com/yuantingapis/go-client",
	}
	args = append(args, files...)
	b, err := exec.Command("protoc", args...).CombinedOutput()
	if err != nil {
		fmt.Printf("%v\n%v\n", string(b), err)
	}
}

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}()

	flag.Parse()
	if *wd == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	*wd, err = filepath.Abs(*wd)
	if err != nil {
		os.Exit(1)
		return
	}

	protoPath := os.Getenv("API_COMMON_PROTOS")
	if protoPath == "" {
		err = errors.New("environment variable API_COMMON_PROTOS not set")
		os.Exit(1)
		return
	}

	genProtoGo(*wd, protoPath)
	genClientGo(*wd, protoPath)
}
