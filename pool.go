package goroutinePool

//协程池
type Pool struct {
	cap,active int   //协程池大小
	ch chan *Goroutine //chan用于保存协程
}

//有可用协程就返回，没有就阻塞，暂不考虑协程池关闭的情况
func (p *Pool)Get()*Goroutine  {
	g := <- p.ch
	p.active++
	return g
}

//协程用完放回
func (p *Pool)Put(g *Goroutine)  {
	p.active--
	g.close()
	p.ch <- g
}

//查看协程总数
func (p *Pool)Cap()int  {
	return p.cap
}

//协程活跃数
func (p *Pool)Active()int  {
	return p.active
}

//获取一个协程池
func NewPool(n int) *Pool {
	pool := Pool{
		cap: n,
		ch:  make(chan *Goroutine,n),
	}

	for i:=0;i<n;i++{
		pool.ch <- &Goroutine{}
	}
	return &pool
}





