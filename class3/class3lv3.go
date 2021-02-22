package main
/*
此文件为测试文件，class3lv3SC.go文件为自建简单通知关闭的Context函数包。


使用管道，自己实现一个可以简单通知关闭的Context函数。

你需要至少实现以下方法

- 初始化方法（如上面的context.WithCancel）
- 通知关闭的方法（如上面的context.Done）

注意这里不是让你完全写的和context标准库一样，只要能实现通知关闭即可。

记得附上使用手册。
*/
import (
	"fmt"
	"myC"
	"time"
)

func worker( ctx myC.MyContext, t *time.Ticker) {
	// Ticker是一个计时器，每过一段时间就向管道Ticker.C中发送一个Time结构体
	go func() {
		defer fmt.Println("worker exit")
		// Using stop channel explicit exit
		for {
			select {
			case <-ctx.Done():
				// ctx.Done()是一个方法，返回一个chan struct{}类型的管道
				// 当主函数中调用了cancel()方法，就会关闭这个管道
				// 按照我们前面讲到，这时会从管道拿到一个零值，用于通知结束
				fmt.Println("recv stop signal")
				return
			case <-t.C:
				fmt.Println("is working")
				// do something
			}
		}
	}()
	return
}

func main() {
	ticker := time.NewTicker(time.Second) // 创建一个计时器

	// 创建一个context，Background()方法返回一个空的context，表示最上层
	// ctx是一个Context类型，cancel是一个func() 类型的函数
	ctx, cancel := myC.InitContext()

	go worker(ctx, ticker)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(50 * time.Millisecond) // 给一定的时间打印信息
}