package main

import (
	"data-structure/linked_list"
	"fmt"
)

func main() {
	head := linked_list.CreateList()
	new_node := linked_list.Node{
		Value: "mojtaba",
		Next: nil,
	}
	

	// new_appended_head := linked_list.AppendToList(&head, &new_node)
	new_prepended_head := linked_list.PrependToList(&head, &new_node)
	fmt.Println(new_prepended_head.Next.Next.Next)
}
