package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int
	var wg sync.WaitGroup

	wg.Add(1) // 每创建一个goroutine，就把任务队列中任务的数量加1
	go func() {
		for i := 0; i < 10; i++ {

			sum = sum + i
		}
		wg.Done() // 任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
	}()

	/* 如果要往队列中添加一个新的任务，wg.Add(1)
	wg.Add(1) // 每创建一个goroutine，就把任务队列中任务的数量加1
	go func() {
		for i := 0; i < 10; i++ {

			sum = sum + i
		}
		wg.Done() // 任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
	}()
	*/

	wg.Wait() // 这里会发生阻塞，直到队列中所有任务完成就会解除阻塞

	fmt.Println(sum)
}

/*
sync.WaitGroup是等待一组协程结束。它实现了一个类似任务队列的结构，你可以向队列中加入任务，任务
完成后就把任务从队列中移除，如果队列中的任务没有全部完成，队列就会触发阻塞以阻止程序继续运行。
sync.WaitGroup只有3个方法，Add()，Done()，Wait()。 其中Done()是Add(-1)的别名。简单的来
说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。
*/
