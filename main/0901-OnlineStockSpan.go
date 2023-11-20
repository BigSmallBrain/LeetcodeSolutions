// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/2 20:01
// -----------------------------------------------
package main

// 单调栈

type StockSpanner struct {
	prices []int
	stack  []int
}

func ConstructorStockSpanner() StockSpanner {
	return StockSpanner{}
}

func (s *StockSpanner) Next(price int) (res int) {

	for len(s.stack) > 0 && price >= s.prices[s.stack[len(s.stack)-1]] {
		s.stack = s.stack[:len(s.stack)-1]
	}

	if len(s.stack) == 0 {
		res = len(s.prices) + 1
	} else {
		res = len(s.prices) - s.stack[len(s.stack)-1]
	}
	s.prices = append(s.prices, price)
	s.stack = append(s.stack, len(s.prices)-1)

	return
}
