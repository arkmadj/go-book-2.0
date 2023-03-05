package memo

type Func func(key string, done <-chan struct{}) (interface{}, error)

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
	done     <-chan struct{}
	response chan<- result
}

type Memo struct {
	requests, cancels chan request
}
