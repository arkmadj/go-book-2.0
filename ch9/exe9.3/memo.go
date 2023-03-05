package memo

type Func func(key string, done <-chan struct{}) (interface{}, error)
