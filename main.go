package main

import (
	// "data-structure/linked_list"
	"data-structure/stack"
	"fmt"
)

func main() {

	//------------------------------------------
	// 				LINKED LIST
	// -----------------------------------------

	// head := linked_list.CreateList()
	// new_node := linked_list.Node{
	// 	Value: "mojtaba",
	// 	Next: nil,
	// }
	

	// new_appended_head := linked_list.AppendToList(&head, &new_node)
	// new_prepended_head := linked_list.PrependToList(&head, &new_node)
	// new_nth_head := linked_list.AddToNthChild(&head, &new_node, 2)
	// fmt.Println(new_nth_head.Next.Next)


	// ------------------------------------------
	// 					STACK
	// ------------------------------------------

	var stack stack.Stack;

	stack.Push("mojtaba")
	stack.Push("rakhisi")
	stack.Push("developer")

	for len(stack) > 0 {
		value , ok := stack.Pop()
		if ok {
			fmt.Println(value)
		}
	} 
}
