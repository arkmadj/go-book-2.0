package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func logCommandError(context string, err error) {
	ee, ok := err.(*exec.ExitError)
	if !ok {
		log.Fatalf("%s: %s", context, err)
	}
	log.Printf("%s: %s", context, err)
	os.Stderr.Write(ee.Stderr)
	os.Exit(1)
}

func packages(patterns []string) []string {
	args := []string{"list", "-f{{.ImportPath}}"}
	args = append(args, patterns...)
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		logCommandError("resolve packages", err)
	}
	return strings.Fields(string(out))
}
