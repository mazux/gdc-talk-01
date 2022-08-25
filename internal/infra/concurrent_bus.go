package infra

import (
	"fmt"
	"math"
	"time"
)

type ConcurrentBus struct {
	bus        *Bus
	cmds       []interface{}
	resultChan chan cmdResult
}

type BackoffStrategy func(int) time.Duration

func (b *ConcurrentBus) Handle() {
	cmdsLen := len(b.cmds)
	for _, cmd := range b.cmds {
		go b.handle(cmd)
	}

	b.cmds = []interface{}{}
	for i := 0; i < cmdsLen; i++ {
		if result := <-b.resultChan; result.err != nil {
			fmt.Println("error: ", result.err)
			b.cmds = append(b.cmds, result.cmd)
		}
	}
}

func (b *ConcurrentBus) Retry(sleep time.Duration) {
	if len(b.cmds) == 0 {
		fmt.Println()
		fmt.Println(time.Now().Format(time.StampMilli), "no retries... ending...")
		return
	}

	time.Sleep(sleep)
	fmt.Println()
	fmt.Println(time.Now().Format(time.StampMilli), "retrying...")
	b.Handle()
}

func (b *ConcurrentBus) RetryWithStrategy(strategy BackoffStrategy, maxRetries int) {
	for i := 0; i < maxRetries; i++ {
		b.Retry(strategy(i))
	}
}

func (b *ConcurrentBus) handle(cmd interface{}) {
	b.resultChan <- cmdResult{cmd, b.bus.Handle(cmd)}
}

func NewConcurrentBus(bus *Bus, cmds []interface{}) *ConcurrentBus {
	resultChan := make(chan cmdResult, len(cmds))
	return &ConcurrentBus{bus, cmds, resultChan}
}

func LinearBackoffStrategy() BackoffStrategy {
	return func(retryCount int) time.Duration {
		return time.Duration(2*retryCount) * time.Second
	}
}

func ExponentialBackoffStrategy() BackoffStrategy {
	return func(retryCount int) time.Duration {
		return time.Duration(math.Pow(float64(2), float64(retryCount))) * time.Second
	}
}

type cmdResult struct {
	cmd interface{}
	err error
}
