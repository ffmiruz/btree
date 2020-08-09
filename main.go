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
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var tree btree
	for _, value := range numbers {
		tree.Insert(value)
		fmt.Println("root:", tree.root.keys)
	}
	fmt.Println(tree.root.keys, tree.root.child[0].keys, tree.root.child[1].keys,
		tree.root.child[1].child[0].keys, tree.root.child[1].child[1].keys)

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

// Insert adds a value into the tree.
func (n *node) Insert(value int) {
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

			top := &node{keys: make([]int, 0, order),
				child: make([]*node, 0, order),
				leaf:  false,
			}
			top.keys = append(top.keys, promoted)
			nn := *n
			top.child = append(top.child, &nn)
			top.child = append(top.child, rnode)
			*n = *top

			fmt.Println(top.keys, nn.keys, rnode.keys)
			return
		}
		return
	}

	n.child[pos].Insert(value)

}
