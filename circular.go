/* 
*
* Vasu Mahalingam Nov 2017
*
*/

package main

import (
	"fmt"
)

// Circular Linked List Implementation
// List contains Integers
// Maintain the circular list in ascending order

// 11------0-----7--8--9--10
// ^ Rear  ^Head   (Arrange in ascending order)

type CircularList struct {
	elem int
	next *CircularList
}

var rear *CircularList

func (list *CircularList) new_element(e int) *CircularList {
	return &CircularList{
		elem: e,
		next: nil,
	}
}

func (list *CircularList) insert(e int) {
	if rear == nil {
		rear = list.new_element(e)
		rear.next = rear
		return
	}
	prev := rear
	curr := rear.next
	for {
		if e <= curr.elem {
			break
		}
		prev = curr
		curr = curr.next
		if curr == rear.next {
			break
		}
	}
	temp := list.new_element(e)
	prev.next = temp
	temp.next = curr
	if e > rear.elem {
		rear = temp
	}
}

func (list *CircularList) print() {
	if rear == nil {
		fmt.Println("List is empty!")
		return
	}
	cur := rear
	for {
		fmt.Println(cur.elem)
		cur = cur.next
		if cur == rear {
			break
		}
	}
}

func (list *CircularList) delete(e int) {
	if rear == nil {
		fmt.Println("Circular list is empty")
		return
	}
	prev := rear
	cur := prev.next
	for {
		if e <= cur.elem {
			break
		}
		prev = cur
		cur = cur.next
		if cur == rear.next {
			break
		}
	}
	if cur.elem != e { // no match
		fmt.Println("Element not found!")
		return
	}
	if cur == prev { // single node 
		rear = nil
		//delete(cur) // GC will reap
		return
	}
	if cur == rear {
		rear = prev // fix rear if last node is deleted
	}
	prev.next = cur.next
	//delete(cur) // GC will reap
	return
}

func main() {
	list := &CircularList{}
	// insert elements into the list
	list.insert(10)
	list.insert(8)
	list.insert(0)
	list.insert(7)
	list.insert(12)
	list.insert(0)
	list.insert(9)
	list.insert(11)
	list.print()
	// Delete inserted elements from the list
	list.delete(12)
	list.delete(0)
	list.delete(9)
	list.delete(11)
	list.delete(13)
	list.print()
	list.delete(10)
	list.delete(0)
	list.delete(7)
	list.delete(8)
	list.print()
}

