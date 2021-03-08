package datastructure

import "fmt"

type List struct {
	Val  int
	Next *List
}

func SearchList(l *List, key int) (cur *List) {
	_, cur = searchList(l, key)
	return
}

func Insert(l *List, key int) {
	temp := *l // change a address, otherwise point to endless
	*l = List{Val: key, Next: &temp}
}

func Delete(l *List, key int) {
	prev, cur := searchList(l, key)
	if cur != nil {
		if prev != nil {
			prev.Next = cur.Next
		} else {
			*l = *(cur.Next)
		}
	}
}

func searchList(l *List, key int) (*List, *List) {
	var prev *List
	for ; l != nil && l.Val != key; l = l.Next {
		prev = l
	}
	return prev, l
}

func (l *List) String() string {
	if l == nil {
		return "nil"
	}
	return fmt.Sprintf("%v %v", l.Val, l.Next)
}

func init() {
	l := &List{Val: 1, Next: nil}
	fmt.Println(l)
	Insert(l, 2)
	Insert(l, 3)
	Insert(l, 4)
	Insert(l, 5)
	fmt.Println(l)
	fmt.Println(SearchList(l, 2))
	Delete(l, 4)
	fmt.Println(l)
}
