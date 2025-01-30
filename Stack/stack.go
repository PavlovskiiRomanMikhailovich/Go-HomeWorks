package main

import "fmt"

type Node struct {
	elem int
	Link *Node
}

type Stack struct {
	Head *Node
}

// Push добавление элемента
func (t *Stack) Push(value int) {
	node := &Node{elem: value, Link: t.Head}
	t.Head = node
}

// PrintAll Вывод всех элементов
func (t *Stack) PrintAll() {
	node := t.Head
	for node.Link != nil {
		fmt.Println(node.elem)
		node = node.Link
	}
	fmt.Println(node.elem)
}

// Pop Удаление первого элемента
func (t *Stack) Pop() {
	if t.IsEmpty() {
		fmt.Println("Error pop: Stack is empty")
	} else {
		node := t.Head
		t.Head = node.Link
	}
}

// IsEmpty Проверка пустой ли стек
func (t *Stack) IsEmpty() bool {
	if t.Head == nil {
		return true
	}
	return false
}

// Len Подсчет длинны
func (t *Stack) Len() int {
	len := 1
	node := t.Head
	for node.Link != nil {
		node = node.Link
		len += 1
	}
	return len
}

func main() {
	test := Stack{Head: &Node{elem: 1, Link: nil}}
	test.Push(2)
	test.Push(3)
	test.Push(4)
	test.Push(5)
	test.PrintAll()
	test.Pop()
	fmt.Println("aaaa")
	test.PrintAll()

}
