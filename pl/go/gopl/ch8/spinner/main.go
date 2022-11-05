// main goroutine 将计算菲波那契数列的第45个元素值
// 由于计算函数使用低效的递归，所以会运行相当长时间，所以用一个可见的标识来表明程序依然在正常运行
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d)=%d", n, fibN)
}

// spinner 一个小动画，表示程序在正常运行
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/ ` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// fib计算斐波那契数列
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

/*
	8.1. Goroutines
	1. 在Go语言中，每一个并发的执行单元叫作一个goroutine。
	2. 当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。
	3. 主函数返回时，所有的goroutine都会被直接打断，程序退出。除了从主函数退出或者直接终止程序之外，
	没有其它的编程方法能够让一个goroutine来打断另一个的执行
	4. 但是之后可以看到一种方式来实现这个目的，通过goroutine之间的通信来让一个goroutine请求其它的goroutine，
	并让被请求的goroutine自行结束执行。
*/
