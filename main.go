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
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	var tree btree
	for _, value := range numbers {
		tree.Insert(value)
	}

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
	t.root.Insert(value)
}

// insert adds the value into the tree.
func (n *node) Insert(value int) {
	pos := sort.SearchInts(n.keys, value)
	if n.leaf {

		// https://github.com/golang/go/wiki/SliceTricks#insert
		n.keys = append(n.keys, 0)
		copy(n.keys[pos+1:], n.keys[pos:])
		n.keys[pos] = value

		fmt.Println(n.keys)
		// if len(n.keys) == order {
		// 	n.split()
		// }
		return
	}
	n.child[pos].Insert(value)

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
