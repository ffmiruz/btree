// B-tree implementation.
// With m as tree order, satisfy the followings:
// 1. all leaf nodes are at the same level.
// 2. all non-leaf nodes (except root) have at least m/2 child at most m child.
// 3. number of keys is one less from the number of children for non-leaf nodes.
// 	  at least m/2 for leaf nodes, at most m-1
// 4. The root may have as few as 2 children unless the tree is the root alone.
//
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
}

// Insert adds a value into the tree.
func (t *btree) Insert(value int) {
	if t.root == nil {
		// Cap == order to leave room before checking for overflow
		t.root = &node{keys: make([]int, 0, order),
			child: make([]*node, 0, order),
			leaf:  true,
		}
		t.root.keys = append(t.root.keys, value)
		return
	}
	k, ch := t.root.Insert(value)

	// Handle root overflow case.
	// Create a single key node from promoted key to be a new root.
	// Previous root as left child, promoted node as right child.
	if ch != nil {
		newRoot := &node{keys: make([]int, 0, order),
			child: make([]*node, 0, order),
			leaf:  false,
		}
		newRoot.keys = append(newRoot.keys, k)
		newRoot.child = append(newRoot.child, t.root)
		newRoot.child = append(newRoot.child, ch)
		t.root = newRoot
	}
}

// Insert adds a value into the node.
func (n *node) Insert(value int) (int, *node) {
	// Find the position to insert the value or the child to follow
	pos := sort.SearchInts(n.keys, value)

	if n.leaf {
		n.keys = append(n.keys, 0)
		copy(n.keys[pos+1:], n.keys[pos:])
		n.keys[pos] = value

		// Split the node
		if len(n.keys) == order {
			mid := order / 2
			promoted := n.keys[mid]

			// Will be the right child node of the promoted key
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
	k, c := n.child[pos].Insert(value)

	// Place returned values into node
	if c != nil {
		n.keys = append(n.keys, 0)
		copy(n.keys[pos+1:], n.keys[pos:])
		n.keys[pos] = k

		// Account for expansion due to new key; update new child insert
		// position by 1 to the right.
		posc := pos + 1
		n.child = append(n.child, c)
		copy(n.child[posc+1:], n.child[posc:])
		n.child[posc] = c
	}

	if len(n.keys) == order {
		mid := order / 2
		promoted := n.keys[mid]

		rnode := &node{keys: make([]int, mid, order),
			child: make([]*node, 0, order),
			leaf:  false,
		}
		copy(rnode.keys, n.keys[mid+1:])
		n.keys = n.keys[:mid]

		// Deal with child
		rnode.child = n.child[mid+1:]
		//// Prevent dropping rightmost child
		n.child = n.child[:mid+1]

		return promoted, rnode
	}
	return 0, nil
}

// Ugly print
func (n *node) print() {
	fmt.Println(n.keys)
	for i := range n.child {
		n.child[i].print()
	}
}

// Todo:
// Handle duplicate key case
