package main

import (
	"bufio"
	"bytes"
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

func ancestors(packageNames []string) []string {
	targets := make(map[string]bool)
	for _, pkg := range packageNames {
		targets[pkg] = true
	}

	args := []string{"list", `-f={{.ImportPath}} {{join .Deps " "}}`, "..."}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		logCommandError("find ancestors", err)
	}
	var pkgs []string
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		fields := strings.Fields(s.Text())
		pkg := fields[0]
		deps := fields[1:]
		for _, dep := range deps {
			if targets[dep] {
				pkgs = append(pkgs, pkg)
				break
			}
		}
	}
	return pkgs
}
