package codestyle

import (
	"fmt"
	"time"
)

type list []string

type Strategy func(t Target)

func (s Strategy) Start(l list){
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
	x0 := func(t Target) {fmt.Printf("attack %v at %v\n", t.l[0],t.n)}
	x1 := func(t Target) {fmt.Printf("attack %v at %v\n", t.l[2],t.n.Add(time.Hour))}
	x2 := func(t Target) {fmt.Printf("attack %v at %v\n", t.l[1],t.n.AddDate(0,0,1))}
	l := list{"gg","mm","bb"}
	Strategy(x0).Start(l)
	Strategy(x1).Start(l)
	Strategy(x2).Start(l)
}
