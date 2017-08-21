package main

import "fmt"

func main(){
	ch := make(chan int)

	go func(){
		var sum int
		for i := 0; i < 10; i++ {
			sum = sum + i
		}
		ch <- sum
	}()

	fmt.Println(<-ch)
}

//无缓冲的通道指的是通道的大小为0，也就是说，这种类型的通道在接收前没有能力保存任何值，
//它要求发送goroutine和接收goroutine同时准备好，才可以完成发送和接收操作。
//
//从上面无缓冲的通道定义来看，发送goroutine和接收gouroutine必须是同步的，
//同时准备后，如果没有同时准备好的话，先执行的操作就会阻塞等待，直到另一个相
//对应的操作准备好为止。这种无缓冲的通道我们也称之为同步通道。
//
//在前面的例子中，我们为了演示goroutine，防止程序提前终止，都是使用sync.WaitGroup
//进行等待，现在的这个例子就不用了，我们使用同步通道来等待。
//
//在计算sum和的goroutine没有执行完，把值赋给ch通道之前，fmt.Println(<-ch)会一直等待
//，所以main主goroutine就不会终止，只有当计算和的goroutine完成后，并且发送到ch通道的操
//作准备好后，同时<-ch就会接收计算好的值，然后打印出来。