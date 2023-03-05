package memo

type Memo struct {
	f     Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)
