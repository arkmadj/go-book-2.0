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

func read(owner, repo, number string) {
	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	body := issue.Body
	if body == "" {
		body = "<empty>\n"
	}
	fmt.Printf("Repo: %s%s\nNumber: %s\nUser: %s\nTitle: %s\n\n%s", owner, repo, number, issue.User.Login, issue.Title, body)
}
