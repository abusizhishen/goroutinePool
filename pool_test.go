package goroutinePool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPool(t *testing.T) {
	var n = 10
	pool := NewPool(n)
	assert.Equal(t,n,pool.cap)
	assert.IsType(t, pool, &Pool{})
}

func TestPool_Get(t *testing.T) {
	pool := NewPool(10)
	g := pool.Get()
	assert.NotNil(t,g)
	assert.IsType(t,g,&Goroutine{})
}

func TestPool_Put(t *testing.T) {
	pool := NewPool(10)
	g := pool.Get()
	g.Do(func() {
		fmt.Println("do a task")
	})

	pool.Put(g)
}
