package utime

import (
	"sync/atomic"
	"time"
)

var (
	statAt = int64(0)
)

func init() {
	n := time.Now().Unix() / 60
	atomic.StoreInt64(&statAt, n)

	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			n := time.Now().Unix() / 60
			atomic.StoreInt64(&statAt, n)
		}
	}()
}

func StatAt() int64 {
	return atomic.LoadInt64(&statAt)
}

func UnixSec() int64 {
	return time.Now().Unix()
}
