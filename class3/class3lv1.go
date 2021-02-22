package main
/*并发求1~50000以内的素数。注意，这里的并发是为了加速程序运行的速度（压榨多核cpu），无意义的并发不得分。*/
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main(){
	var a[50001] bool
	re := make([]int, 0)
	for i:= 2; i <= 50000; i++ {
		if !a[i] {
			wg.Add(1)
			go func(x int) {//将x的倍数标记为非素数
				defer wg.Done()
				for j:= x + x; j < 50000; j += x {
					a[j] = true
				}
			}(i)
		}

	}
	wg.Wait()
	for i:= 1; i <= 50000; i++{
		if !a[i] {
			re = append(re, i)
		}
	}
	fmt.Println(re)
}
