package list

type List struct {
	length int
	Head   *Node
	Tail   *Node
}

func NewList() *List {
	l := &List{}
	l.length = 0

	return l
}

func (l *List) Length() int {
	return l.length
}

func (l *List) IsEmpty() bool {
	return l.length == 0
}

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

func (l *List) Prepend(value interface{}) {
	node := NewNode(value)
	if l.Length() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		//formerhead points to currentHead
		formerHead := l.Head
		//formerHead previous is now the node that is being added
		formerHead.Prev = node

		//next from new node points to former head
		node.Next = formerHead
		//head from list is now  the new node
		l.Head = node
	}
	l.length++
}

func (l *List) Append(value interface{}) {
	node := NewNode(value)
	if l.Length() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		//formerTail points to current Tail
		formerTail := l.Tail
		//formerTail next points to new node
		formerTail.Next = node

		//new node previous points to former tail
		node.Prev = formerTail
		//new tail of list is new node
		l.Tail = node
	}
	l.length++
}
