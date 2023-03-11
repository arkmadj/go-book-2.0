package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func mainExample() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}
}
