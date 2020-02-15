package codestyle

import (
	"fmt"
	"time"
)

type Attacker interface {
	attacking()
}

type List []string

type Strategy func(t *Target)

type Target struct {
	l []string
	n time.Time
	Strategy
}

func (t *Target) attacking(){
	t.Strategy(t)
}

func (s Strategy) Start(l List){
	t := &Target{
		l:        l,
		n:        time.Now(),
		Strategy: s,
	}
	t.attacking()
}

func StrategyUsage() {
	s0 := func(t *Target) {fmt.Printf("%v will attack at %v\n", t.l[0],t.n)}
	s1 := func(t *Target) {fmt.Printf("%v will attack at %v\n", t.l[2],t.n.Add(time.Hour))}
	s2 := func(t *Target) {fmt.Printf("%v will attack at %v\n", t.l[1],t.n.AddDate(0,0,1))}
	l := List{"iron-man","spider-man","black-widow"}
	Strategy(s0).Start(l)
	Strategy(s1).Start(l)
	Strategy(s2).Start(l)
}
