package goroutinePool

//协程结构体
type Goroutine struct {}

//协程执行某些任务
func (g *Goroutine) Do(fun func()) {
	fun()
}

//关闭协程池时可能有一些资源回收相关的
func (g *Goroutine)close()  {}

