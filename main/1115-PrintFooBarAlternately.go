// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/8/14 10:04
// -----------------------------------------------
package main

type FooBar struct {
	n   int
	ch1 chan bool
	ch2 chan bool
}

func NewFooBar(n int) *FooBar {
	f := &FooBar{
		n:   n,
		ch1: make(chan bool, 1),
		ch2: make(chan bool),
	}
	f.ch1 <- true
	return f
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch1
		printFoo()
		fb.ch2 <- true
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch2
		printBar()
		fb.ch1 <- true
	}
}
