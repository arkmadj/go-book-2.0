package memo

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Close() { close(memo.requests) }

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}
