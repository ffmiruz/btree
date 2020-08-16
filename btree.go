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
)

type Btree struct {
	Root *node
}

type node struct {
	key   []int
	child []*node
	leaf  bool
}

const order int = 5

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
		12, 13, 14, 15, 16, 17}
	var tree Btree
	for _, value := range numbers {
		tree.Insert(value)
	}
	tree.Root.Print()
}

// Insert adds a value into the tree.
func (t *Btree) Insert(value int) {
	if t.Root == nil {
		// Cap == order to leave room before checking for overflow
		t.Root = &node{key: make([]int, 0, order+10),
			child: make([]*node, 0, order),
			leaf:  true,
		}
		t.Root.key = append(t.Root.key, value)
		return
	}
	k, c := t.Root.insert(value)

	// Handle Root overflow case.
	// Create a single key node from promoted key to be a new Root.
	// Previous Root as left child, promoted node as right child.
	if c != nil {
		newRoot := &node{key: make([]int, 0, order),
			child: make([]*node, 0, order),
			leaf:  false,
		}
		newRoot.key = append(newRoot.key, k)
		newRoot.child = append(newRoot.child, t.Root)
		newRoot.child = append(newRoot.child, c)
		t.Root = newRoot
	}
}

// Insert adds a value into the node.
func (n *node) insert(value int) (int, *node) {
	// Find the position to insert the value or the child to follow
	pos := cutMark(n.key, value)

	if n.leaf {
		// Slot value into corrent position.
		// From https://github.com/golang/go/wiki/SliceTricks#insert
		n.key = append(n.key, 0)
		copy(n.key[pos+1:], n.key[pos:])
		n.key[pos] = value

		// Split the node
		if len(n.key) == order {
			mid := order / 2
			promoted := n.key[mid]

			// Will be the right child node of the promoted key
			rnode := &node{key: make([]int, mid, order),
				leaf: true,
			}
			rnode.key = n.key[mid+1:]
			n.key = n.key[:mid]
			return promoted, rnode
		}
		return 0, nil
	}
	k, c := n.child[pos].insert(value)

	// Place returned values into node
	if c != nil {
		n.key = append(n.key, 0)
		copy(n.key[pos+1:], n.key[pos:])
		n.key[pos] = k

		// Account for expansion due to new key; update new child insert
		// position by 1 to the right.
		posc := pos + 1
		n.child = append(n.child, c)
		copy(n.child[posc+1:], n.child[posc:])
		n.child[posc] = c
	}

	// Split the node
	if len(n.key) == order {
		mid := order / 2
		promoted := n.key[mid]

		rnode := &node{key: make([]int, mid, order),
			leaf: false,
		}
		rnode.key = n.key[mid+1:]
		n.key = n.key[:mid]

		rnode.child = n.child[mid+1:]
		//// Prevent dropping rightmost child
		n.child = n.child[:mid+1]

		return promoted, rnode
	}
	return 0, nil
}

// Ugly print
func (n *node) Print() {
	fmt.Println(n.key, n.leaf, len(n.child))
	for i := range n.child {
		n.child[i].Print()
	}
}

// Search position to insert an integer into an ascending order sorted slice.
// For this implementation, duplicate value will slot behind existing value.
func cutMark(sorted []int, v int) int {
	pos := 0
	for i := range sorted {
		if v > sorted[i] {
			pos = i + 1
			continue
		}
		break
	}
	return pos
}
