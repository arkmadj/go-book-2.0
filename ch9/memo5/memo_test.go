package memo_test

import (
	"testing"

	memo "github.com/ahmad/go-book-2.0/ch9/memo5"
	"github.com/ahmad/go-book-2.0/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}
