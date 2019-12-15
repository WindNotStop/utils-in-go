package channel

import (
	"fmt"
	"time"
)

func OrChannel() {
	//模拟一个可关闭done的协程
	done := func(after time.Duration) <-chan interface{} {
		done := make(chan interface{})
		go func() {
			defer close(done)
			//模拟实际操作
			time.Sleep(after)
		}()
		return done
	}
	//如果有一个子done关闭则关闭
	var or func(done ...<-chan interface{}) <-chan interface{}
	or = func(done ...<-chan interface{}) <-chan interface{} {
		switch len(done) {
		case 0:
			return nil
		case 1:
			return done[0]
		}
		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			defer fmt.Println("orDone")
			switch len(done) {
			case 2:
				select {
				case <-done[0]:
				case <-done[1]:
				}
			default:
				select {
				case <-done[0]:
				case <-done[1]:
				case <-done[2]:
				case <-or(append(done[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
	start := time.Now()
	<-or(done(time.Second), done(2*time.Second), done(3*time.Second), done(4*time.Second), done(5*time.Second), done(6*time.Second))
	fmt.Println("done after", time.Since(start))
}
