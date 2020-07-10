package linkedlists

type ListNode struct {
	value string
	next *ListNode
}

type LinkedList struct{
	count int
	head *ListNode
	tail *ListNode
}

func create(value string) *ListNode{
	return &ListNode{value, nil}
}

func (l *LinkedList) getNode(index int) (*ListNode, error){
	current := l.head
	for i:=0; i<index; i++{
		current = current.next
	}
	return current, nil
}

func (l *LinkedList) IsEmpty() bool{
	return l.Size()==0
}

func (l *LinkedList) Size() int{
	return l.count
}

func (l *LinkedList) InsertAt(index int, value string) error{
	element := create(value)
	if l.IsEmpty(){
		l.count ++
		l.head = element
		l.tail = element
		return nil
	}
	if index == 0{
		l.count ++
		element.next = l.head
		l.head = element
		element = nil
		return nil
	}
	if index == l.count{
		l.tail.next = element
		l.tail = element
		l.count ++
		element = nil
		return nil
	}
	current, err := l.getNode(index-1)
	if err != nil{
		return err
	}
	element.next = current.next
	current.next = element
	current = nil
	element = nil
	l.count ++
	return nil
}

func (l *LinkedList) GetHead() *ListNode{ return l.head}
func (l *LinkedList) GetTail() *ListNode{ return l.tail}

func (l *LinkedList) GetValueAt(index int) *string{
	current, err := l.getNode(index)
	if err != nil{
		return nil
	}
	return &current.value
}

func (l *LinkedList) RemoveAt(index int) (*string, error){
	if l.count == 1{
		value := l.head.value
		l.head = nil
		l.tail = nil
		l.count--
		return &value, nil
	}
	if index == 0{
		current := l.head
		l.head = current.next
		current.next =nil
		value := current.value
		l.count--
		current = nil
		return &value, nil
	}

	current, err := l.getNode(index - 1)
	if err !=  nil{
		return nil, err
	}

	value := current.next.value
	current.next = current.next.next
	if index == l.count -1 {
		l.tail = current
	}
	l.count--
	current = nil
	return &value, nil
}

func (l *LinkedList) Clear(){
	current := l.head
	for current != nil{
		temp := current
		current = current.next
		temp.next = nil
		temp = nil
	}
	current = nil
	l.count = 0
	l.head = nil
	l.tail = nil
}

func (l *LinkedList) GetValue() []string{
	values := make([]string, 0, l.Size())
	current := l.head
	for current != nil{
		values = append(values, current.value)
		current = current.next
	}
	current = nil
	return values
}

func (l *LinkedList) GetIndexOf(value string) int{
	current := l.head
	index := 0
	for current != nil{
		if value == current.value{
			return index
		}
		current = current.next
		index ++
	}
	current = nil
	return -1
}
