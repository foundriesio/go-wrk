package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)


var (
	checkFileFilterIn = []string{
		"./objects/",
		"./deltas/",
	}
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	repoDir := flag.String("repo", cwd, "A path to an ostree repo")
	outList := flag.String("out", cwd, "A path to a file to output object list to")
	flag.Parse()
	log.Printf("Repo dir: %s\n", *repoDir)

	of, err := os.Create(*outList)
	if err != nil {
		log.Fatalf("Failed to create a file to output object list to: %s\n", err.Error())
	}
	defer of.Close()

	if err := filepath.Walk(*repoDir, func(fullPath string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			log.Fatalf("Failed to walk through a repo: %s\n", walkErr.Error())
		}
		if info.IsDir() {
			return nil
		}
		relPath := strings.Replace(fullPath, *repoDir, ".", 1)
		if !filterRepoFiles(relPath, checkFileFilterIn) {
			return nil
		}
		log.Printf("%s\n", relPath)
		_, err = of.WriteString(fmt.Sprintf("%s\n", relPath[2:]))
		if err != nil {
			log.Fatalf("Failed to write to a file: %s\n", err.Error())
		}
		return nil
	}); err != nil {
		log.Fatalf("Failed to walk through a repo directory: %s\n", err.Error())
	}
}

func filterRepoFiles(path string, filter []string) bool {
	for _, f := range filter {
		if strings.HasPrefix(path, f) {
			return true
		}
	}
	return false
}