package main

import (
	"fmt"
	"sync/atomic"
)

// 链接 http://www.jianshu.com/p/d2fc1c998fc9

type singleton struct {
}

// private
var instance *singleton

// ---------------------------------------------------------------------------
// public
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}		// not thread safe
	}
	return instance
}
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
var mu sync.Mutex

func GetInstance() *singleton {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &singleton{}		// unnecessary locking if instance already created
	}
	return instance
}

/**
这里使用了Go的sync.Mutex,其工作模型类似于Linux内核的futex对象，具体实现极其简单，性能也有保证
初始化时填入的0值将mutex设定在未锁定状态，同时保证时间开销最小
这一特性允许将mutex作为其它对象的子对象使用
 */
// ---------------------------------------------------------------------------


// ---------------------------------------------------------------------------
var mu sync.Mutex

func GetInstance() *singleton {
	if instance == nil {  	// not yet perfect. since it's not fully atomic
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

/**
这是一个不错的方法，但是还不完美，因为编译器优化没有检查实例存储状态
如果使用sync/atomic包的话，就可以自动加载和设置标记
 */
// ---------------------------------------------------------------------------



// ---------------------------------------------------------------------------
// atomic方式
var initialized uint32
func GetInstance() *singleton {
	if atomic.LoadInt32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if initialized == 0{
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}


// ----------------------------------------------------------
// 比较好的一种方式
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
// ---------------------------------------------------------------------------

func main() {
	i := GetInstance()
	fmt.Println(i)
}