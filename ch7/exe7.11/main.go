package main

import "sync"

type PriceDB struct {
	sync.Mutex
	dp map[string]int
}
