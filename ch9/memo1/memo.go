package memo

type Memo struct {
	f     Func
	cache map[string]result
}
