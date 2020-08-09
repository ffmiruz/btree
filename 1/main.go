package main

import (
	"fmt"
	"sort"
)

type node struct {
	keys  []int
	child []*node
	leaf  bool
}

type btree struct {
	root *node
}

const order int = 5

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var tree btree
	for _, value := range numbers {
		tree.Insert(value)
	}
	fmt.Println("root:", tree.root)

}

// Insert adds a value into the tree.
func (t *btree) Insert(value int) {
	if t.root == nil {
		// cap order to leave room before checking for overflow
		t.root = &node{keys: make([]int, 0, order),
			child: make([]*node, 0, order),
			leaf:  true,
		}
		t.root.keys = append(t.root.keys, value)

		return
	}
	new, promoted := t.root.Insert(value)
	fmt.Println("promo: ", promoted)
	if new != nil {
		t.root.leaf = false

		pos := sort.SearchInts(t.root.keys, promoted)

		t.root.keys = append(t.root.keys, 0)
		copy(t.root.keys[pos+1:], t.root.keys[pos:])
		t.root.keys[pos] = promoted

		fmt.Println(t.root.keys, t.root.child, "new keys")

		// posn := pos + 1
		// var empty node
		// t.root.child = append(t.root.child, &empty)
		// copy(t.root.child[posn+1:], t.root.child[posn:])
		// t.root.child[posn] = new

	}
}

// insert adds the value into the tree.
func (n *node) Insert(value int) (*node, int) {
	pos := sort.SearchInts(n.keys, value)
	if n.leaf {

		// https://github.com/golang/go/wiki/SliceTricks#insert
		n.keys = append(n.keys, 0)
		copy(n.keys[pos+1:], n.keys[pos:])
		n.keys[pos] = value

		if len(n.keys) == order {
			mid := order / 2
			promoted := n.keys[mid]

			rnode := &node{keys: make([]int, mid, order),
				child: make([]*node, 0, order),
				leaf:  true,
			}
			copy(rnode.keys, n.keys[mid+1:])

			n.keys = n.keys[:mid]

			fmt.Println(n.keys, rnode.keys, "xxxx")

			return rnode, promoted

		}
		return nil, 0
	}
	new, promoted := n.child[pos].Insert(value)
	fmt.Println(n.keys, "bbbbb")

	if new != nil {
		if len(n.child) < 1 {
			newTop := &node{keys: make([]int, 1, order),
				child: make([]*node, 2, order),
				leaf:  false,
			}
			newTop.keys[0] = promoted
			newTop.child = append(newTop.child, n)
			newTop.child = append(newTop.child, new)
		}
		n.leaf = false

		posn := pos + 1

		n.keys = append(n.keys, 0)
		copy(n.keys[pos+1:], n.keys[pos:])
		n.keys[pos] = promoted

		var empty node
		n.child = append(n.child, &empty)
		copy(n.child[posn+1:], n.child[posn:])
		n.child[posn] = new
	}

	return nil, 0

	// if len(n.keys) < order && n.leaf {
	// 	n.keys = append(n.keys, value)
	// 	sort.Ints(n.keys)
	// 	return
	// }
	// if !n.leaf {

	// }
	// for i := 0; i < len(n.keys); i++ {
	// 	switch {
	// 	case value <= n.keys[i]:
	// 		if n.left == nil {
	// 			n.left = &node{value: value}
	// 			return
	// 		}
	// 		n.left.Insert(value)
	// 	case value > n.value:
	// 		if n.right == nil {
	// 			n.right = &node{value: value}
	// 			return
	// 		}
	// 		n.right.Insert(value)
	// 	}
	// }

	// if len(n.keys) > order-1 {
	// 	split()
	// }

}

func slot(s []int, i, value int) {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = value
}
