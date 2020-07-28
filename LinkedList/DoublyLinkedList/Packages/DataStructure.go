package doublyLinkedList

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{value, nil, nil}
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) Prev() *Node {
	return node.prev
}

func (node *Node) Next() *Node {
	return node.next
}
