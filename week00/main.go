package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("Your GOPATH has not been set!")
	}
	fmt.Println("GOPATH\t", gopath)

	gobin := filepath.Join(gopath, "bin")
	path := os.Getenv("PATH")

	if !strings.Contains(path, gobin) {
		log.Fatalf("Your PATH does not contain %s", gobin)
	}

	fmt.Println("GOBIN\t", gobin)
	fmt.Println("Success!")

}