package pool

import (
	"bytes"
	"sync"
)

// BytesPool is a singleton Buffer factory
// strings.Builder is not suitable as Reset() is inefficient
var BytesPool = NewBufferPool()

type Buffer struct {
	*bytes.Buffer
	pool Pool
}

func (s *Buffer) Free() {
	s.pool.put(s)
}

// A Pool is sync.Pool wrapper
type Pool struct {
	p *sync.Pool
}

func NewBufferPool() Pool {
	return Pool{
		p: &sync.Pool{
			New: func() interface{} {
				sb := new(bytes.Buffer)
				sb.Grow(20)
				return &Buffer{
					Buffer: sb,
				}
			},
		}}
}

func (p Pool) Get() *Buffer {
	sb := p.p.Get().(*Buffer)
	sb.Reset()
	sb.pool = p
	return sb
}

func (p Pool) put(s *Buffer) {
	p.p.Put(s)
}
