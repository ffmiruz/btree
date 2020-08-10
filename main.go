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
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
		12, 13, 14, 15, 16, 17}
	var tree btree
	for _, value := range numbers {
		tree.Insert(value)
	}
	tree.root.print()

	// fmt.Println(tree.root.child[0].keys, tree.root.child[1].keys)
	// fmt.Println(tree.root.child[0].child[0].keys,
	// 	tree.root.child[0].child[1].keys, //tree.root.child[0].child[2].keys,
	// 	tree.root.child[1].child[0].keys, tree.root.child[1].child[1].keys,
	// 	tree.root.child[1].child[2].keys,
	// )

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
	v, ch := t.root.Insert(value)
	if ch != nil {
		newRoot := &node{keys: make([]int, 0, order),
			child: make([]*node, 0, order),
			leaf:  false,
		}
		newRoot.keys = append(newRoot.keys, v)
		newRoot.child = append(newRoot.child, t.root)
		newRoot.child = append(newRoot.child, ch)

		t.root = newRoot
	}
}

// Insert adds a value into the tree.
func (n *node) Insert(value int) (int, *node) {
	pos := sort.SearchInts(n.keys, value)

	if n.leaf {
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
			return promoted, rnode
		}
		return 0, nil
	}
	v, ch := n.child[pos].Insert(value)
	if ch != nil {
		n.child = append(n.child, ch)
		n.keys = append(n.keys, v)
	}

	if len(n.keys) == order {
		mid := order / 2
		promoted := n.keys[mid]

		rnode := &node{keys: make([]int, mid, order),
			child: make([]*node, 0, order),
			leaf:  false,
		}
		copy(rnode.keys, n.keys[mid+1:])
		fmt.Println("--", n.keys)
		fmt.Println(">>", n.child[2].keys)
		n.keys = n.keys[:mid]

		// deal with child
		rnode.child = n.child[mid+1:]
		// prevent dropping rightmost child
		n.child = n.child[:mid+1]
		fmt.Println("---", n.keys, rnode.keys)

		return promoted, rnode
	}

	return 0, nil

}

func (n *node) print() {
	fmt.Println(n.keys)
	for i := range n.child {
		n.child[i].print()
	}
}
