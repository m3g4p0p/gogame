//go:build ignore

package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	SRC_DIR  = "public"
	DEST_DIR = "dist"
)

func copyFile(srcPath, destPath string) error {
	var srcFile *os.File
	var destFile *os.File
	var err error

	if srcFile, err = os.Open(srcPath); err != nil {
		return err
	}
	defer srcFile.Close()

	if destFile, err = os.Create(destPath); err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}

func copyPublicTree() error {
	return filepath.WalkDir(SRC_DIR, func(path string, d fs.DirEntry, err error) error {
		var destPath string

		if err != nil {
			return err
		}

		if relPath, err := filepath.Rel(SRC_DIR, path); err != nil {
			return err
		} else {
			destPath = filepath.Join(DEST_DIR, relPath)
		}

		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		return copyFile(path, destPath)
	})
}

func copyWasmExec() error {
	cmd := exec.Command("go", "env", "GOROOT")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	goroot := strings.TrimSpace(string(output))
	wasmExecPath := filepath.Join(goroot, "lib", "wasm", "wasm_exec.js")
	wasmDestPath := filepath.Join(DEST_DIR, "wasm_exec.js")

	return copyFile(wasmExecPath, wasmDestPath)
}

func buildMainWasm() error {
	cmd := exec.Command("go", "build", "-o", filepath.Join(DEST_DIR, "main.wasm"), ".")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	if err := os.MkdirAll(DEST_DIR, os.ModePerm); err != nil {
		panic(err)
	}

	if err := buildMainWasm(); err != nil {
		panic(err)
	}

	if err := copyWasmExec(); err != nil {
		panic(err)
	}

	if err := copyPublicTree(); err != nil {
		panic(err)
	}

	fmt.Println("Build completed successfully!")
}
