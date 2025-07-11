package homework4

type node struct {
	key   int
	val   int
	left  *node
	right *node
}
type OrderedMap struct {
	root *node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (n *node) insert(key int, val int, m *OrderedMap) *node {
	if n == nil {
		m.size++
		return &node{key: key, val: val}
	}
	if key < n.key {
		n.left = n.left.insert(key, val, m)
	} else if key > n.key {
		n.right = n.right.insert(key, val, m)
	} else {
		n.val = val
	}
	return n
}
func (m *OrderedMap) Insert(key, value int) {
	m.root = m.root.insert(key, value, m)
}

func (n *node) findMin() *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *node) erase(key int, m *OrderedMap) *node {
	if n == nil {
		return nil
	} else if key < n.key {
		n.left = n.left.erase(key, m)
	} else if key > n.key {
		n.right = n.right.erase(key, m)
	} else {
		m.size--

		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		replace := n.right.findMin()
		n.key = replace.key
		n.val = replace.val
		n.right = n.right.erase(replace.key, m)
	}
	return n
}
func (m *OrderedMap) Erase(key int) {
	m.root = m.root.erase(key, m)
}

func (n *node) containsKey(key int) bool {
	if n == nil {
		return false
	} else if key == n.key {
		return true
	} else if key < n.key {
		return n.left.containsKey(key)
	} else {
		return n.right.containsKey(key)
	}
}

func (m *OrderedMap) Contains(key int) bool {
	if m.root == nil {
		return false
	}
	return m.root.containsKey(key)
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (n *node) foreach(action func(int, int)) {
	if n != nil {
		n.left.foreach(action)
		action(n.key, n.val)
		n.right.foreach(action)
	}
}
func (m *OrderedMap) ForEach(action func(int, int)) {
	m.root.foreach(action)
}
