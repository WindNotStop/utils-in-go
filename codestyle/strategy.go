package codestyle

import "fmt"

type Attacker interface {
	prepare()
	attacking()
	cleanup()
}

type Scheduler struct {
}

func (p *Scheduler) prepare() {
	fmt.Println("prepare")
}

func (p *Scheduler) attacking() {
	fmt.Println("attacking")
}

func (p *Scheduler) cleanup() {
	fmt.Println("cleanup")
}

type Strategy func(Attacker)

func (s Strategy) Start(a Attacker) {
	s(a)
}

func StrategyUsage() {
	p := &Scheduler{}
	s0 := func(a Attacker) {
		a.attacking()
		a.prepare()
		a.cleanup()
	}
	s1 := func(a Attacker) {
		a.prepare()
		a.attacking()
		a.cleanup()
	}
	s2 := func(a Attacker) {
		a.attacking()
		a.cleanup()
		a.prepare()
	}
	Strategy(s0).Start(p)
	Strategy(s1).Start(p)
	Strategy(s2).Start(p)
}
