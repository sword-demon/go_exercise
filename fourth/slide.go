package fourth

import (
	"sync"
	"time"
)

// 请求次数
var countLockSlide = 0

// 单次请求
type slideItem struct {
	time  time.Time // 请求时间
	allow bool      // 是否允许
}

/**
	滑动窗口计数器
*/

type CounterSlide struct {
	rate  int           //计数周期内最多允许的请求数
	cycle time.Duration //计数周期

	itemMap map[int]slideItem
	lock    sync.Mutex
}

func (l *CounterSlide) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	countLockSlide++

	tmp := slideItem{
		time: time.Now(),
	}
	if l.checkItem(tmp.time) {
		tmp.allow = true
	} else {
		tmp.allow = false
	}
	l.itemMap[countLockSlide] = tmp
	return tmp.allow
}

// 处理map元素
func (l *CounterSlide) dealMap(now time.Time) {
	for key, value := range l.itemMap {
		thisTime := value.time
		if now.Sub(thisTime) > l.cycle {
			delete(l.itemMap, key)
		}
	}
}

// 检查itemArr 中的元素
// 这里也可以只记录允许通过的记录 用切片记录即可
func (l *CounterSlide) checkItem(now time.Time) bool {
	var limit = 0
	for key, value := range l.itemMap {
		thisTime := value.time
		if now.Sub(thisTime) > l.cycle {
			delete(l.itemMap, key)
			continue
		}
		// 允许
		if value.allow {
			limit++
		}
	}

	if limit < l.rate {
		return true
	}
	return false
}

func (l *CounterSlide) Set(r int, cycle time.Duration) {
	if r <= 0 || cycle <= 0 {
		panic("参数异常")
	}

	l.rate = r
	l.cycle = cycle

	l.itemMap = make(map[int]slideItem)
}
