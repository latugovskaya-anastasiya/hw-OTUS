package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	self := &ListItem{}

	self.Value = v
	self.Prev = nil

	if l.first != nil {
		self.Next = l.first
		l.first.Prev = self
	}

	if l.last == nil {
		l.last = self
	}

	l.first = self
	l.len++
	return self
}

func (l *list) PushBack(v interface{}) *ListItem {
	self := &ListItem{}
	if l.first == nil {
		l.first = self
	}
	if l.last == nil {
		l.last = self
	}
	self.Value = v
	self.Next = nil
	self.Prev = l.last

	l.last.Next = self
	l.last = self
	l.len++
	return self
}

func (l *list) Remove(i *ListItem) {
	if i == l.first && l.first.Next == nil {
		l.first = nil
		l.last = nil
		l.len = 0
		return
	}
	if i == l.first && l.first.Next != nil {
		i.Next.Prev = nil
		l.first = i.Next
		l.len--
		return
	}
	if i == l.last && i.Prev != nil {
		i.Prev.Next = nil
		l.len--
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	i.Prev = nil
	i.Next = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.first {
		return
	}

	i.Prev.Next = i.Next
	if i == l.last {
		l.last = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.first.Prev = i

	i.Next = l.first
	i.Prev = nil
	l.first = i
}

/*NewList ...*/
func NewList() List {
	return new(list)
}
