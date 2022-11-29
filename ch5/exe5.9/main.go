package main

import "regexp"

var pattern = regexp.MustCompile(`\$\w+|\${\w+}`)
