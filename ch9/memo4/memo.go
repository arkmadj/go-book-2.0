package memo

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}
