package main

import (
	"os"
	"fmt"
	"bufio"
	"sort"
)

// all IO operations are defined here
// printf("%d %s\n", &a, &str)
// ==========================================================================
// IO
// ==========================================================================
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
	writer.Flush()
}
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }
// Don't use scanLine()
// func scanLine() string {
// 	s, _ := reader.ReadString('\n')
// 	return s
// }
// ==========================================================================

// Vector is defined here
// := vec()
// for i:=0; i<3; i++ {
// .pushback(i)
// }
// printf("%v\n",)
// num :=.popback()
// printf("%v, popped: %v\n",, num)
//.sort(func(i, j int) bool {
// 	return[i].(int) <[j].(int)
// })
// ==========================================================================
// Vector 
// ==========================================================================
type vector []interface{}
func vec() vector {
	v := make(vector, 0)
	return v
}
func (v *vector) empty() bool {
	return v.len() == 0
}
func (v *vector) len() int {
	return len((*v))
}
func (v *vector) pushback(x interface{}) {
	(*v) = append((*v), x)
}
func (v *vector) pushfront(x interface{}) {
	(*v) = append((*v), x)
}
func (v *vector) popback() (interface{}, bool) {
	length := len((*v))
	if length > 0 {
		ret := (*v)[length-1]
		(*v) = (*v)[:length - 1]
		return ret, true
	}
	return 0, false
}
func (v *vector) popfront() (interface{}, bool) {
	length := len((*v))
	if length > 0 {
		ret := (*v)[0]
		(*v) = (*v)[1:]
		return ret, true
	}
	return 0, false
}
func (v *vector) first() (interface{}, bool) {
	length := len((*v))
	if length > 0 {
		return (*v)[0], true
	}
	return 0, false
}
func (v *vector) last() (interface{}, bool) {
	length := len((*v))
	if length > 0 {
		return (*v)[length-1], true
	}
	return 0, false
}
func (v *vector) get(pos int) (interface{}, bool) {
	length := len((*v))
	if length > 0 && 0 <= pos && pos < length {
		return (*v)[pos], true
	}
	return 0, false
}
func (v *vector) insert(pos int, x interface{}) {
	(*v) = append((*v)[:pos+1], (*v)[pos:]...)
	(*v)[pos] = x
}
func (v *vector) remove(pos int) {
	(*v) = append((*v)[:pos], (*v)[pos+1:]...)
}
func (v *vector) sort(f func(int, int) bool) {
	sort.Slice((*v), f)
}
// =========================================================================


// =========================================================================
// Priority Queue 
// =========================================================================
type priorityQueue []interface{}
func pq() priorityQueue {
	p := make([]interface{}, 0)
	return p
}
func (q *priorityQueue) empty() bool {
	return q.len() == 0
}
func (q *priorityQueue) len() int {
	return len((*q))
}
func (q *priorityQueue) swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}
func (q *priorityQueue) push(f func(int, int) bool, num int) {
	(*q) = append((*q), num)
	parent := (q.len() - 2) / 2
	q.heapify(f, parent)
}
func (q *priorityQueue) pop(f func(int, int) bool) (interface{}, bool) {
	if q.len() > 0 {
		highest := (*q)[0] 
		q.swap(0, q.len() - 1)
		(*q) = (*q)[:q.len() - 1]
		q.heapify(f, 0)
		return highest, true
	}
	return 0, false
}
func (q *priorityQueue) heapify(f func(int, int) bool, curr int) {
	largest := curr
	length := q.len()
	left := 2 * curr + 1
	right := 2 * curr + 2
	if left < length && f(left, largest) {
		largest = left
	}
	if right < length && f(right, largest) {
		largest = right
	}
	if largest != curr {
		q.swap(largest, curr)
		q.heapify(f, largest)
	}
}
func (q *priorityQueue) peek() (interface{}, bool) {
	if q.len() > 0 {
		return (*q)[0], true
	}
	return 0, false
}
// =========================================================================

// =========================================================================
// Binary tree
// =========================================================================
type Node struct {
	val int
	left *Node
	right *Node
}

func node(val int) *Node {
	return &Node{
		val: val,
		left: nil,
		right: nil,
	}
}
// =========================================================================
func main() {
	// a, b := 0, 0
	// c := ""
	// x := 0
	// printf("Enter numbers: \n")
	// scanf("%d %d\n", &a, &b)
	// scanf("%d %s\n", &x, &c)
	// printf("%d %s\n", a+b+x, c)

	// arr := vec()
	// for i:=0; i<3; i++ {
	// 	arr.pushback(i)
	// }
	// printf("%v\n", arr)

	// num, _ := arr.popback()
	// printf("%v, popped: %v\n", arr, num)

	// arr.pushback(34)
	// arr.pushback(23)
	// arr.pushback(12)
	// arr.pushback(84)
	// arr.insert(2, 90)

	// arr.sort(func(i, j int) bool {
	// 	return arr[i].(int) < arr[j].(int)
	// })

	// printf("sorted: %v\n", arr)

	// heap := pq()
	// fn := func(i, j int) bool {
	// 	return heap[i].(int) > heap[j].(int)
	// }
	// heap.push(fn, 1)
	// printf("%v\n", heap)
	// heap.push(fn, 3)
	// printf("%v\n", heap)
	// heap.push(fn, 9)
	// printf("%v\n", heap)
	// heap.push(fn, 2)
	// printf("%v\n", heap)
	// heap.push(fn, 7)
	// printf("%v\n", heap)

	// top, _ := heap.peek()
	// printf("topmost: %v\n", top)
	
	// n, _ := heap.pop(fn)
	// printf("%v\n", n)
	// n, _ = heap.pop(fn)
	// printf("%v\n", n)
	// n, _ = heap.pop(fn)
	// printf("%v\n", n)
	// n, _ = heap.pop(fn)
	// printf("%v\n", n)
	// n, _ = heap.pop(fn)
	// printf("%v\n", n)

	tc := 0
	scanf("%d\n", &tc)
	for t := 1; t <= tc; t++ {
		// Here accept inputs
		
		// Here do something

		printf("Case #%d: \n", t)
	}
}
