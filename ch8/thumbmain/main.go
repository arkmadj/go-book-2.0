package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ahmad/go-book-2.0/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// thumb, err := thumbnail.ImageFile(input.Text())
		thumb, err := thumbnail.ImageFile(input.Text())
		// thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
