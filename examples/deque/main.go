package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/comparator"
	"github.com/liyue201/gostl/containers/deque"
)

func main() {
	q := deque.New(0)
	q.PushBack(2)
	q.PushFront(1)
	q.PushBack(3)
	q.PushFront(4)
	fmt.Printf("%v\n", q)

	sort.Sort(q.Begin(), q.End(), comparator.BuiltinTypeComparator)
	fmt.Printf("%v\n", q)
}