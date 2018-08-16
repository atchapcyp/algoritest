package main

import "fmt"

type stack struct {
	pile [10]int
	top  int
}

func (x *stack) push(v int) {
	x.pile[x.top] = v
	x.top++
}

func (x *stack) pop() int {
	v := x.pile[x.top-1]
	x.pile[x.top-1] = 0
	x.top--
	println("in pop : v = ", v)
	return v
}

func starthanoi() {
	var s1, s2, s3 stack
	// fmt.Printf("begin s1 is : %v ", s1)
	// println()
	// s1.push(10)
	// fmt.Printf("Push 10 s1 is : %v ", s1)
	// s1.push(9)
	// fmt.Printf("Push 9 s1 is  : %v ", s1)
	// println()
	// s1.pop()
	// fmt.Printf("default top s1 after pop is : %v ", s1)
	// s1.pop()
	// println()
	// fmt.Printf("default top s1 after pop is : %v ", s1)
	// println()
	s1.push(5)
	s1.push(4)
	s1.push(3)
	s1.push(2)
	s1.push(1)
	fmt.Printf("begin s1 is : %v \n", s1)
	fmt.Printf("begin s2 is : %v \n", s2)
	fmt.Printf("begin s3 is : %v \n", s3)
	move(&s1, &s3, &s2, 5)
	fmt.Printf("begin s1 is : %v \n", s1)
	fmt.Printf("begin s2 is : %v \n", s2)
	fmt.Printf("begin s3 is : %v \n", s3)
}

func move(s1, s3, s2 *stack, n int) {
	fmt.Printf(" s1 is : %v \n", s1)
	fmt.Printf(" s2 is : %v \n", s2)
	fmt.Printf(" s3 is : %v \n\n\n", s3)
	if n == 1 {
		disk := s1.pop()
		s3.push(disk)
		return
	}
	if n == 2 {
		disk := s1.pop()
		s2.push(disk)
		disk = s1.pop()
		s3.push(disk)
		disk = s2.pop()
		s3.push(disk)
	}
	if n > 2 {
		move(s1, s2, s3, n-1)
		move(s1, s3, s2, 1)
		move(s2, s3, s1, n-1)
	}

}
