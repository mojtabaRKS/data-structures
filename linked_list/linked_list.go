package linked_list

type Node struct {
	Value interface{}
	Next *Node
}

func CreateList() Node {
	fourth_and_last_node := Node{
		Value: "fourth node",
		Next: nil,
	}

	third_node := Node{
		Value: "third node",
		Next: &fourth_and_last_node,
	}

	second_node := Node{
		Value: "second node",
		Next: &third_node,
	}

	head := Node{
		Value: "it is the first node",
		Next: &second_node,
	}


	return head
}

func AppendToList(head, node *Node) Node {
	node.Next = head
	return *node
}

func PrependToList(head, node *Node) Node {
	current := head
	
	for current != node {
		
		if current.Next == nil {
			current.Next = node
			break
		}

		current = current.Next
	}

	return *head
}

func AddToNthChild(head, node *Node, lvl int) Node {
	
	current := head

	i := 1

	for i <= lvl {
		if i == lvl {
			node.Next = current.Next
			current.Next = node
			break
		}

		current = current.Next
		i++
	}

	return *head
}

