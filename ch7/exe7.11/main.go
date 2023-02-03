package main

import (
	"net/http"
	"strconv"
	"sync"
)

type PriceDB struct {
	sync.Mutex
	d map[string]int
}

func (p *PriceDB) Create(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "No integer price given", http.StatusBadGateway)
		return
	}

	if _, ok := p.db[item]; ok {

	}
}
