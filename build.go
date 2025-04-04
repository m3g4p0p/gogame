//go:build ignore

package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	SRC_DIR  = "public"
	DEST_DIR = "dist"
)

func copyTree() error {
	return filepath.WalkDir(SRC_DIR, func(path string, d fs.DirEntry, err error) error {
		var destPath string
		var srcFile *os.File
		var destFile *os.File

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

		if srcFile, err = os.Open(path); err != nil {
			return err
		} else {
			defer srcFile.Close()
		}

		if destFile, err = os.Create(destPath); err != nil {
			return err
		} else {
			defer destFile.Close()
		}

		_, err = io.Copy(destFile, srcFile)
		return err
	})
}

func main() {
	// Ensure the dist directory exists
	os.MkdirAll(DEST_DIR, os.ModePerm)

	// Build the Go WebAssembly binary
	cmd := exec.Command("go", "build", "-o", filepath.Join(DEST_DIR, "main.wasm"), ".")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	if err := copyTree(); err != nil {
		panic(err)
	}

	fmt.Println("Build completed successfully!")
}
