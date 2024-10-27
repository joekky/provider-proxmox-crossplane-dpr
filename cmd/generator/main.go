/*
Copyright 2021 Upbound Inc.
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/pkg/pipeline"
	"github.com/joekky/provider-proxmox-crossplane/config"
)

func main() {
	if len(os.Args) < 3 {
		panic("source and version arguments are required")
	}

	// Get current working directory as root
	rootDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("cannot get current working directory: %v", err))
	}

	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}

	pipeline.Run(config.GetProvider(), absRootDir)
}
