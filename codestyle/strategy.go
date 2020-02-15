package codestyle

import (
	"fmt"
	"time"
)

type List []string

type Strategy func(t Target)

func (s Strategy) Start(l List){
	t := Target{
		l:        l,
		n:        time.Now(),
	}
	s(t)
}

type Target struct {
	l []string
	n time.Time
}

func StrategyUsage() {
	s0 := func(t Target) {fmt.Printf("attack %v at %v\n", t.l[0],t.n)}
	s1 := func(t Target) {fmt.Printf("attack %v at %v\n", t.l[2],t.n.Add(time.Hour))}
	s2 := func(t Target) {fmt.Printf("attack %v at %v\n", t.l[1],t.n.AddDate(0,0,1))}
	l := List{"gg","mm","bb"}
	Strategy(s0).Start(l)
	Strategy(s1).Start(l)
	Strategy(s2).Start(l)
}
