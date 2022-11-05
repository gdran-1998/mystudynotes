# 8. Goroutines和Channels

并发程序：同时进行多个人物的程序。Go语言的并发程序用两种手段来实现：

- CSP模式，用 goroutine 和 channel 方式来实现，值在不同的运行实例（goroutine）中传递。
- 多线程共享内存，通过一些同步机制，比如锁之类的机制来实现并发。

## 8.1. Goroutines

每一个并发的执行单元叫作一个 goroutine 。

当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

主函数返回时，所有的goroutine都会被直接打断，程序退出。除了从主函数退出或者直接终止程序之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行。但是之后可以看到一种方式来实现这个目的，通过goroutine之间的通信来让一个goroutine请求其它的goroutine，并让被请求的goroutine自行结束执行。

spinner例子(spinning和菲波那契的计算。分别在独立的函数中，但两个函数会同时执行。)

```go
func main() {
    go spinner(100 * time.Millisecond)
    const n = 45
    fibN := fib(n) // slow
    fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
    for {
        for _, r := range `-\|/` {
            fmt.Printf("\r%c", r)
            time.Sleep(delay)
        }
    }
}

func fib(x int) int {
    if x < 2 {
        return x
    }
    return fib(x-1) + fib(x-2)
}
```

## 8.4. Channels

一个channel是一个通信机制，它让一个goroutine通过它给另一个goroutine发送值信息。int 类型的地channel写成 chan int 。

内置的make函数，创建channel：

```go
ch := make(chan int)
```

一个channel对应一个make创建的底层数据结构的引用。

channel有发送和接送操作，发送语句将值从goroutine通过channel发送到执行接收操作的goroutine。

```go
ch <- x
x = <- ch
<- ch
```

对关闭的channl再作发送操作将panic；已经被close过的channel依旧可以接收之前发送成功的数据；channel中没有数据的话将产生零值。

有缓冲和无缓冲channel：

```go
ch = make(chan int)
ch = make(chan int,0)
ch = make(chan int,4)
```

### 8.4.1 无缓冲channel

对于一个无缓冲的channel，光发送不接收将导致发送者阻塞，直到有goroutine来接收，反之亦然。（同步操作）



### 8.4.2 串联的channel（pipeline）

















