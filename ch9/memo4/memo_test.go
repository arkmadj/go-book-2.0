package memo_test

import (
	"testing"

	memo "github.com/ahmad/go-book-2.0/ch9/memo4"
	"github.com/ahmad/go-book-2.0/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
