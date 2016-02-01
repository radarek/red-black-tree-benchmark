package red_black_tree

type Color int

const (
	Red Color = iota
	Black
)

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	key    int
	color  Color
}

var NilNode = NewNilNode()

func NewNilNode() *Node {
	node := &Node{color: Black}
	return node
}

func NewNode(key int, color Color) *Node {
	node := &Node{key: key, color: color, left: NilNode, right: NilNode, parent: NilNode}
	return node
}

func (self *Node) isRed() bool {
	return self.color == Red
}

func (self *Node) isBlack() bool {
	return self.color == Black
}

func (self *Node) isNil() bool {
	return self == NilNode
}

type RedBlackTree struct {
	Root *Node
	size int
}

func NewRedBlackTree() *RedBlackTree {
	tree := &RedBlackTree{Root: NilNode, size: 0}
	return tree
}

func (tree *RedBlackTree) Add(key int) {
	tree.insert(NewNode(key, Red))
}

func (tree *RedBlackTree) insert(x *Node) {
	tree.insertHelper(x)

	x.color = Red
	for x != tree.Root && x.parent.color == Red {
		if x.parent == x.parent.parent.left {
			y := x.parent.parent.right
			if !y.isNil() && y.color == Red {
				x.parent.color = Black
				y.color = Black
				x.parent.parent.color = Red
				x = x.parent.parent
			} else {
				if x == x.parent.right {
					x = x.parent
					tree.leftRotate(x)
				}
				x.parent.color = Black
				x.parent.parent.color = Red
				tree.rightRotate(x.parent.parent)
			}
		} else {
			y := x.parent.parent.left
			if !y.isNil() && y.color == Red {
				x.parent.color = Black
				y.color = Black
				x.parent.parent.color = Red
				x = x.parent.parent
			} else {
				if x == x.parent.left {
					x = x.parent
					tree.rightRotate(x)
				}
				x.parent.color = Black
				x.parent.parent.color = Red
				tree.leftRotate(x.parent.parent)
			}
		}
	}
	tree.Root.color = Black
}

func (tree *RedBlackTree) insertHelper(z *Node) {
	y := NilNode
	x := tree.Root
	for !x.isNil() {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y.isNil() {
		tree.Root = z
	} else {
		if z.key < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}
	tree.size += 1
}

func (tree *RedBlackTree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if !y.left.isNil() {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent.isNil() {
		tree.Root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.left = x
	x.parent = y
}

func (tree *RedBlackTree) rightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if !y.right.isNil() {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent.isNil() {
		tree.Root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.right = x
	x.parent = y
}

func (tree *RedBlackTree) Minimum(x *Node) *Node {
	if x == nil {
		x = tree.Root
	}
	for !x.left.isNil() {
		x = x.left
	}
	return x
}

func (tree *RedBlackTree) Maximum(x *Node) *Node {
	if x == nil {
		x = tree.Root
	}
	for !x.right.isNil() {
		x = x.right
	}
	return x
}

func (tree *RedBlackTree) successor(x *Node) *Node {
	if !x.right.isNil() {
		return tree.Minimum(x.right)
	}
	y := x.parent
	for !y.isNil() && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (tree *RedBlackTree) predecessor(x *Node) *Node {
	if !x.left.isNil() {
		return tree.Maximum(x.left)
	}
	y := x.parent
	for !y.isNil() && x == y.left {
		x = y
		y = y.parent
	}
	return y
}

func (tree *RedBlackTree) Delete(z *Node) *Node {
	var x, y *Node
	if z.left.isNil() || z.right.isNil() {
		y = z
	} else {
		y = tree.successor(z)
	}
	if y.left.isNil() {
		x = y.right
	} else {
		x = y.left
	}
	x.parent = y.parent

	if y.parent.isNil() {
		tree.Root = x
	} else {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	}

	if y != z {
		z.key = y.key
	}

	if y.color == Black {
		tree.deleteFixup(x)
	}

	tree.size -= 1
	return y
}

func (tree *RedBlackTree) deleteFixup(x *Node) {
	for x != tree.Root && x.color == Black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				tree.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == Black && w.right.color == Black {
				w.color = Red
				x = x.parent
			} else {
				if w.right.color == Black {
					w.left.color = Black
					w.color = Red
					tree.rightRotate(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = Black
				w.right.color = Black
				tree.leftRotate(x.parent)
				x = tree.Root
			}
		} else {
			w := x.parent.left
			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				tree.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.right.color == Black && w.left.color == Black {
				w.color = Red
				x = x.parent
			} else {
				if w.left.color == Black {
					w.right.color = Black
					w.color = Red
					tree.leftRotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = Black
				w.left.color = Black
				tree.rightRotate(x.parent)
				x = tree.Root
			}
		}
	}
	x.color = Black
}

func (tree *RedBlackTree) Search(key int) *Node {
	x := tree.Root
	for !x.isNil() && x.key != key {
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (tree *RedBlackTree) InorderWalk(callback func(int)) {
	x := tree.Minimum(nil)
	for !x.isNil() {
		callback(x.key)
		x = tree.successor(x)
	}
}

func (tree *RedBlackTree) ReverseInorderWalk(callback func(int)) {
	x := tree.Maximum(nil)
	for !x.isNil() {
		callback(x.key)
		x = tree.predecessor(x)
	}
}
