package memo

import "sync"

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}
