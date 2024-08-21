package main

import (
	"fmt"
	"errors"
)

type node struct {
	key string
	val interface{}
	prev *node
	next *node
}

type DList struct {
	begin *node
}

type DL struct {
	capacity int
	num int
	items map[string]*node
	dlist *DList
}


func min(a, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}


func NewDlist() *DList {
	return &DList {
		begin: nil,
	}
}

func (dlist *DList) MoveToBegin(node *node) {
	// move it to beginning of list
	tmp := dlist.begin
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.next = tmp
	tmp.prev = node
	node.prev = nil
	dlist.begin = node
}


func (dlist *DList) Insert(node *node) error {
	if node == nil {
		return errors.New("Null node cannot be inserted")
	}
	if dlist.begin == nil {
		dlist.begin = node
	} else {
		node.next = dlist.begin
		dlist.begin.prev = node
		node.prev = nil
		dlist.begin = node
	}
	return nil
}


func (dlist *DList) Delete(node *node) error {
	if node == nil ||  dlist.begin == nil {
		return errors.New("Null node cannot be deleted")
	}

	tmp := dlist.begin.next
	if node == dlist.begin {
		dlist.begin = tmp
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}
	

	return nil
}


func NewNode(key string, val interface{}) *node {
	return &node{
		key: key,
		val: val,
		prev: nil,
		next: nil,
	}
}





func NewDL(capacity int) *DL {
	return &DL{
		capacity: capacity,
		dlist: NewDlist(),
		items: make(map[string]*node, capacity),
		num: 0,
	}
}

func (dl *DL) Get(key string) interface{} {
	if _, ok := dl.items[key]; ok  { 
		lastptr := dl.items[key]
		dl.dlist.MoveToBegin(lastptr)
		return lastptr.val
	}
	return nil
}

func (dl *DL) Put(key string, val interface{}) error {
	if len(dl.items) > dl.capacity {
		return errors.New("Capacity reached")
	}
	dl.Delete(key)
	node := NewNode(key, val)
	if node == nil {
		return errors.New("Could not allocate node")
	}
	dl.dlist.Insert(node)
	dl.items[key] = node
	dl.num++
	return nil
}

func (dl *DL) Delete(key string) error {
	if _, ok := dl.items[key]; ok {
		dl.dlist.Delete(dl.items[key])
		delete(dl.items, key)
		dl.num--
		return nil
	} else {
		return errors.New("key not found")
	}
}

func (dl *DL) Last() string {
	if dl.dlist.begin == nil {
		return ""
	}
	return dl.dlist.begin.key
}

func main() {
	dl := NewDL(10)
	fmt.Println(dl)

	dl.Put("madhukar", 1)
	dl.Put("taru", 2)
	dl.Put("myra", 3)
	dl.Put("vihaan", 4)

	fmt.Println(dl)

	last := dl.Last()

	fmt.Println(last)

	dl.Get("madhukar")

	last = dl.Last()
	fmt.Println(last)

	dl.Delete("madhukar")

	last = dl.Last()

	fmt.Println(last)

	dl.Get("myra")


	last = dl.Last()

	fmt.Println(last)

	fmt.Println(dl.num)

}