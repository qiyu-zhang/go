package myC
/*自建简单通知关闭的Context函数包*/

//包含管道的简单结构体
type MyContext struct {
	ch chan struct{}
}

//调用结构体中的管道
func (c MyContext) Done ()  chan struct{}{
	return c.ch
}
//初始化生成myContext结构体和相应管道撤销函数
func InitContext() (c MyContext, cancel func()) {
	c = MyContext{make(chan struct{})}
	return c, func() {
		close(c.ch)
	}

}