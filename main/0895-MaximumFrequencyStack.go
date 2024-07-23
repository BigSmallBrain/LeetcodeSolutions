// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/23 20:45
// -----------------------------------------------
package main

// 哈希表

type FreqStack struct {
	group   map[int][]int
	freq    map[int]int
	maxFreq int
}

func FreqStackConstructor() FreqStack {
	return FreqStack{map[int][]int{}, make(map[int]int), 0}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	this.group[this.freq[val]] = append(this.group[this.freq[val]], val)
	this.maxFreq = max(this.maxFreq, this.freq[val])
}

func (this *FreqStack) Pop() int {
	val := this.group[this.maxFreq][len(this.group[this.maxFreq])-1]
	this.group[this.maxFreq] = this.group[this.maxFreq][:len(this.group[this.maxFreq])-1]
	this.freq[val]--
	if len(this.group[this.maxFreq]) == 0 {
		this.maxFreq--
	}
	return val
}
