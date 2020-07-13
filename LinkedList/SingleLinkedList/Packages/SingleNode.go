package SingleLinkedList

type Node struct {
	value interface{}
	next  *Node //鏈表的節點
}

//新建一個節點指向先前的節點，返回這個節點的地址
func NewNode(data interface{}) *Node {
	// fmt.Println(data)
	return &Node{data, nil}
}

//返回值
func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) Next() *Node {
	return node.next
}
