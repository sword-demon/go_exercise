package fourth

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test_SlidingWindowLimiter_Run(t *testing.T) {
	// qps=100 windowNum = 10
	limiter := NewSlidingWindowLimiter(100*time.Millisecond, 1*time.Second, 100)

	total := 500
	i := 1
	handler := func(expectSucc, expectFail int64) {
		var wg sync.WaitGroup
		var succ, fail int64
		// so many request coming at the same time
		for j := 0; j < total; j++ {
			go func() {
				wg.Add(1)
				err := limiter.Run(func() {
					time.Sleep(300 * time.Millisecond)
				})
				if err != nil {
					atomic.AddInt64(&fail, 1)
				} else {
					atomic.AddInt64(&succ, 1)
				}
				wg.Done()
			}()
		}

		wg.Wait()
		if succ != expectSucc || fail != expectFail {
			t.Errorf("i:%d expect succ: %d fail:%d, actual succ:%d fail:%d\n", i, expectSucc, expectFail, succ, fail)
		}
	}
	go handler(100, 400)
	time.Sleep(time.Second)
}
