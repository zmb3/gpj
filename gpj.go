package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

// GopathLibs gets all of the library packages in $GOPATH.
func GopathLibs() []*build.Package {
	return gopathPackages(false)
}

// GopathPackages gets all packages in $GOPATH, including
// libraries and commands.
func GopathPackages() []*build.Package {
	return gopathPackages(true)
}

func gopathPackages(commands bool) []*build.Package {
	var pkgs []*build.Package
	for _, dir := range filepath.SplitList(build.Default.GOPATH) {
		src := filepath.Join(dir, "src")
		if fi, err := os.Stat(src); err != nil || !fi.IsDir() {
			continue
		}
		filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil || !info.IsDir() {
				return nil
			}
			n := info.Name()
			if strings.HasPrefix(n, ".") || n == "testdata" || n == "internal" || n == "testfiles" {
				return filepath.SkipDir
			}
			pkg, err := build.Default.ImportDir(path, 0)
			if err == nil {
				if commands || (!commands && !pkg.IsCommand()) {
					pkgs = append(pkgs, pkg)
				}
			}
			return nil
		})
	}
	return pkgs
}

func main() {
	// a map to track which packages are imported
	imported := make(map[string]bool)

	// start out with all packages unimported
	libs := GopathLibs()
	for _, lib := range libs {
		imported[lib.ImportPath] = false
	}

	// mark imported packages
	pkgs := GopathPackages()
	for _, pkg := range pkgs {
		for _, ip := range pkg.Imports {
			if used, ok := imported[ip]; ok && !used {
				imported[ip] = true
			}
		}
	}

	for pkg, used := range imported {
		if !used {
			fmt.Println(pkg)
		}
	}
}
