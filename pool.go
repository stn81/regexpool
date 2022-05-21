package regexpool

import (
	"regexp"
	"sync"
)

// Pool defines the regex pool
type Pool struct {
	pool    *sync.Pool
	Pattern string
}

// New create a new regex pool for pattern
func New(pattern string) *Pool {
	_ = regexp.MustCompile(pattern)

	return &Pool{
		Pattern: pattern,
		pool: &sync.Pool{
			New: func() interface{} {
				return regexp.MustCompile(pattern)
			},
		},
	}
}

// GetMatcher return a matcher
func (p *Pool) GetMatcher() (m *regexp.Regexp) {
	return p.pool.Get().(*regexp.Regexp)
}

// PutMatcher release the matcher
func (p *Pool) PutMatcher(m *regexp.Regexp) {
	p.pool.Put(m)
}
