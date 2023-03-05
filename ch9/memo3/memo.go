package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}
