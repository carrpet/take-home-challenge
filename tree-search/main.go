package main

import (
	"cmp"
	"slices"
)

// A Tree is a binary tree node containing an id integer
// and left, right pointers to its left and right children.
type Tree struct {
	id          int
	left, right *Tree
}

// A queue is a helper data structure for the breadth-first level labelling
// routine recordLevels.
type queue []levelNode

func enqueue(queue queue, element levelNode) queue {
	queue = append(queue, element)
	return queue
}

func dequeue(queue queue) (levelNode, queue) {
	element := queue[0]
	if len(queue) == 1 {
		return element, queue[:0]
	}

	return element, queue[1:]

}

// A levelNode wraps a Tree node with its level
// in the tree.
type levelNode struct {
	node  *Tree
	level int
}

// A valueLevel represents the value of the id of a Tree node
// and level that it appears in the tree. It is used in favor
// of the intermediate LevelNode representation to facilitate
// testing.
type valueLevel struct {
	value int
	level int
}

// A NullableInt is a boxed int.
type NullableInt *int

func NewNullableInt(val int) NullableInt {
	heapInt := new(int)
	*heapInt = val
	return heapInt

}

// recordLevels traverses the nodes of the tree and returns the
// representation of the tree as a slice of LevelNodes.
func recordLevels(tree *Tree) []valueLevel {
	if tree == nil {
		return []valueLevel{}
	}

	queue := make([]levelNode, 0)
	result := make([]valueLevel, 0)
	queue = enqueue(queue, levelNode{tree, 0})
	var ln levelNode
	for len(queue) > 0 {
		ln, queue = dequeue(queue)
		result = append(result, valueLevel{ln.node.id, ln.level})
		if ln.node.left != nil {
			queue = enqueue(queue, levelNode{ln.node.left, ln.level + 1})
		}
		if ln.node.right != nil {
			queue = enqueue(queue, levelNode{ln.node.right, ln.level + 1})
		}
	}

	return result

}

// scanDuplicates returns the duplicate values and the associated minimum level
// of the input slice vl.
func scanDuplicates(vl []valueLevel) []valueLevel {
	result := []valueLevel{}
	var toAdd *valueLevel
	i := 0
	for i < len(vl)-1 {
		toCompare := vl[i]
		j := i + 1
		for j < len(vl) && vl[j].value == toCompare.value {
			if toAdd == nil {
				toAdd = &valueLevel{toCompare.value, toCompare.level}
			} else {
				if vl[j].level < toAdd.level {
					toAdd.level = vl[j].level
				}

			}
			j++
		}
		if toAdd != nil {
			result = append(result, *toAdd)
			i = j + 1
			toAdd = nil
		} else {
			i = j
		}
	}
	return result
}

func CheckDuplicateIDs(tree *Tree) (value *int, level int) {
	if tree == nil {
		return nil, 0
	}
	valueLevels := recordLevels(tree)

	// sort the valueLevels slice by the entry's value
	slices.SortFunc(valueLevels, func(a, b valueLevel) int {
		return cmp.Compare(a.value, b.value)
	})

	duplicates := scanDuplicates(valueLevels)

	// sort the duplicates by ascending level so the first
	// element has the minimum level
	slices.SortFunc(duplicates, func(a, b valueLevel) int {
		return cmp.Compare(a.level, b.level)
	})

	if len(duplicates) == 0 {
		return nil, 0
	}

	return NewNullableInt(duplicates[0].value), duplicates[0].level

}
