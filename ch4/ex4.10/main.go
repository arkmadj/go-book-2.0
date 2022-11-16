package main

import (
	"fmt"
	"log"
)

func search(query []string) {
	result, err := SearchISuues(query)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		format := "#%-5d %9.9s %.55s\n"
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}
}
